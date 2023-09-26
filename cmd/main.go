package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/inits/config"
	"github.com/freer4an/portfolio-website/inits/mongodb"
	"github.com/freer4an/portfolio-website/internal/api/admin"
	"github.com/freer4an/portfolio-website/internal/api/client"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/freer4an/portfolio-website/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err)
	}
}

func run() error {
	// get env variables
	config, err := config.InitConfig("./configs")
	if err != nil {
		return err
	}

	// logging mode
	if config.Env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()

	client, err := initMongoDB(ctx, config)
	if err != nil {
		return err
	}

	store := repository.New(client,
		config.Database.Name,
		config.Database.CollProject)
	server := server.NewServer(config)
	initAPIs(store, config, server)

	go func() {
		if err := server.Start(config.App.Addr); err != nil {
			log.Fatal().Err(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTSTP)
	<-sig
	gracefullShotdown(ctx, client, server)
	return nil
}

func gracefullShotdown(ctx context.Context, client *mongo.Client, server *server.Server) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("Mongo client closed")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("Server gracefully shutdown")
}

func initMongoDB(ctx context.Context, config *config.Config) (*mongo.Client, error) {
	client, err := mongodb.Connect(ctx, config.Database.Uri)
	if err != nil {
		return nil, err
	}
	if err := mongodb.MongoMigrate(client, config.Database.Name, config.Database.CollProject); err != nil {
		return nil, err
	}
	return client, nil
}

func initAPIs(store *repository.Repository, config *config.Config, server *server.Server) {
	adminAPI := admin.New(store)
	clientAPI := client.New(store)
	r := chi.NewRouter()
	r.Use(helpers.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Mount("/", clientAPI.Routes())
	r.Mount("/admin", adminAPI.Routes())
	server.InitRoutes(r)
	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/freer4an/portfolio-website",
			Intro:       "Portfolio-website generated docs.",
		}))
	}
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if strings.HasPrefix(origin, "http://") || strings.HasPrefix(origin, "https://") {
		// if origin == "http://localhost:8000/" {
		// 	return true
		// }
		return true
	}
	return false
}
