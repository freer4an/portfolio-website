package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	TagName string `bson:"tag_name" json:"tag_name"`
	TagType string `bson:"tag_type" json:"tag_type"`
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Abstract    string             `bson:"abstract" json:"abstract"`
	Description string             `bson:"description,omitempty" json:"description"`
	Url         string             `bson:"url,omitempty" json:"url"`
	Tags        []Tag              `bson:"tags,omitempty" json:"tags"`
	IsFinished  bool               `bson:"finished,omitempty" json:"finished"`
}
