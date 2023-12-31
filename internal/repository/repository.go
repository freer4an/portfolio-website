package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Project ProjectI
}

func New(client *mongo.Client, dbName, coll_name string) *Repository {
	db := client.Database(dbName)
	coll_Projects := db.Collection(coll_name)

	return &Repository{
		Project: NewProjectR(db, coll_Projects),
	}
}
