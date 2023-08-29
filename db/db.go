package db

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStore(client *mongo.Client, dbName, collName string) *Store {
	db := client.Database(dbName)
	coll := db.Collection(collName)
	return &Store{db: db, collection: coll}
}
