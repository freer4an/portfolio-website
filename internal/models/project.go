package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name" validate:"required"`
	Abstract    string             `bson:"abstract" json:"abstract"`
	Description string             `bson:"description,omitempty" json:"description"`
	Url         string             `bson:"url,omitempty" json:"link" validate:"url"`
	Tags        []Tag              `bson:"tags,omitempty" json:"tags"`
	IsFinished  bool               `bson:"finished,omitempty" json:"finished" validate:"boolean"`
}
