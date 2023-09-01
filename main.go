package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freer4an/portfolio-website/db"
	"github.com/freer4an/portfolio-website/server"
	"github.com/freer4an/portfolio-website/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config := util.InitConfig(".")

	ctx := context.Background()
	client := runMongo(ctx, config)
	if err := mongoMigrate(ctx, config, client); err != nil {
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
}

func runMongo(ctx context.Context, config util.Config) *mongo.Client {
	if config.DBuri == "" {
		log.Fatal().Msg("Empty addres")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBuri))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to MongoDB")
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal().Err(err).Msg("failed to ping MongoDB client")
	}

	return client
}

func mongoMigrate(ctx context.Context, config util.Config, client *mongo.Client) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	db := client.Database(config.DBname)
	coll := db.Collection(config.CollName)
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Info().Msgf("Created index {%s} for collection {%s} of db {%s}", name, coll.Name(), db.Name())
	return nil
}