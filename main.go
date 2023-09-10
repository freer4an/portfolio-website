package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freer4an/portfolio-website/db"
	"github.com/freer4an/portfolio-website/db/mongodb"
	"github.com/freer4an/portfolio-website/server"
	"github.com/freer4an/portfolio-website/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config := util.InitConfig(".")

	// logging mode
	if config.Env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()

	client := mongodb.MongoClient(ctx, config)
	if err := mongodb.MongoMigrate(ctx, config, client); err != nil {
		log.Fatal().Err(err)
	}
	store := db.NewStore(client, config.DBname, config.CollName)
	server := server.NewServer(ctx, config, store)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTSTP)
	<-sig
	gracefullShotdown(ctx, client, server)
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
