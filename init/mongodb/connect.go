package mongodb

import (
	"context"
	"fmt"

	"github.com/freer4an/portfolio-website/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect and get mongo client
func MongoClient(ctx context.Context, config *util.Config) (*mongo.Client, error) {
	if config.DBuri == "" {
		return nil, fmt.Errorf("Empty mongodb uri")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBuri))
	if err != nil {
		return nil, fmt.Errorf("failed connection to MongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Empty mongodb uri: %v", err)
	}

	return client, nil
}
