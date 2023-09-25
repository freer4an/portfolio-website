package mongodb

import (
	"context"
	"fmt"

	"github.com/freer4an/portfolio-website/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// make migrations and add unique index to db 'projects' field 'name'
func MongoMigrate(ctx context.Context, config *util.Config, client *mongo.Client) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	db := client.Database(config.DBname)
	coll := db.Collection(config.CollName)
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return fmt.Errorf("Failed to create index: %v", err)
	}
	log.Info().Msgf("Created index {%s} for collection {%s} of db {%s}", name, coll.Name(), db.Name())
	return nil
}
