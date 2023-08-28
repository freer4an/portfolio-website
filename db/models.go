package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description,omitempty"`
	Url         string             `bson:"url,omitempty"`
	IsFinished  bool               `bson:"finished,omitempty"`
}
