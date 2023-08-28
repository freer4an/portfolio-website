package db

import "go.mongodb.org/mongo-driver/mongo"

const (
	projects = "projects"
)

type Store struct {
	db *mongo.Database
}

func NewStore(client *mongo.Client, dbName string) *Store {
	db := client.Database(dbName)
	return &Store{db: db}
}
