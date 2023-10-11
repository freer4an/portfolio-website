package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect and get mongo client
func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if uri == "" {
		return nil, fmt.Errorf("Empty mongodb uri")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed connection to MongoDB: %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf(": %v", err)
	}

	return client, nil
}
