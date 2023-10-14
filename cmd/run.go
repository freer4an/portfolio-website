package main

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freer4an/portfolio-website/inits/config"
	"github.com/freer4an/portfolio-website/inits/mongodb"
	"github.com/freer4an/portfolio-website/internal/api"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/freer4an/portfolio-website/server"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func run() error {
	// get .yaml variables and environment
	config, err := config.InitConfig("./configs")
	if err != nil {
		return err
	}

	if err = godotenv.Load(); err != nil {
		return fmt.Errorf("Error reading .env %v", err)
	}

	// logging mode
	if config.Env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.TODO()

	client, err := initMongoDB(ctx, config)
	if err != nil {
		return err
	}

	store := repository.New(client,
		config.Database.Name,
		config.Database.CollProject)

	api := api.New(store, temp)
	server := server.New()
	server.BuildRoutes("/projects", api.Project.Routes())
	server.BuildRoutes("/", api.Client.Routes())

	go func() {
		if err := server.Start(config.App.Addr); err != nil {
			log.Fatal().Err(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTSTP)
	<-sig
	gracefullShutdown(client, server)
	return nil
}

func gracefullShutdown(client *mongo.Client, server *server.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("./front/templates/*.html"))
}
