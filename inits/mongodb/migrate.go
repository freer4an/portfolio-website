package mongodb

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// make migrations and add unique index to db 'projects' field 'name'
func MongoMigrate(client *mongo.Client, db_name, collection string) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	db := client.Database(db_name)
	coll := db.Collection(collection)
	name, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return fmt.Errorf("Failed to create index: %v", err)
	}
	log.Info().Msgf("Created index {%s} for collection {%s} of db {%s}", name, coll.Name(), db.Name())
	return nil
}
