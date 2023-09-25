package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/init/mongodb"
	"github.com/freer4an/portfolio-website/internal/api/admin"
	"github.com/freer4an/portfolio-website/internal/api/client"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/freer4an/portfolio-website/server"
	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
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
	config, err := util.InitConfig(".")
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
		config.DBname,
		config.CollName)
	server := server.NewServer(config)
	initAPIs(store, config, server)

	go func() {
		if err := server.Start(config.HttpAddrSite); err != nil {
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

func initMongoDB(ctx context.Context, config *util.Config) (*mongo.Client, error) {
	client, err := mongodb.MongoClient(ctx, config)
	if err != nil {
		return nil, err
	}
	if err := mongodb.MongoMigrate(ctx, config, client); err != nil {
		return nil, err
	}
	return client, nil
}

func initAPIs(store *repository.Repository, config *util.Config, server *server.Server) {
	adminAPI := admin.New(config, store)
	clientAPI := client.New(store)
	r := chi.NewRouter()
	r.Use(helpers.Logger)
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
