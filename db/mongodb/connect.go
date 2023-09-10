package mongodb

import (
	"context"

	"github.com/freer4an/portfolio-website/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect and get mongo client
func MongoClient(ctx context.Context, config util.Config) *mongo.Client {
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
