package repository

import (
	"context"

	"github.com/freer4an/portfolio-website/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagI interface {
	AddToProject(ctx context.Context, project_name string, tags ...models.Tag) error
	DeleteFromProject(ctx context.Context, project_name string, tags ...string) error
}

type TagsR struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTagsR(db *mongo.Database, collection *mongo.Collection) *TagsR {
	return &TagsR{
		db:         db,
		collection: collection,
	}
}

func (repo *TagsR) AddToProject(ctx context.Context, project_name string, tags ...models.Tag) error {
	filter := bson.M{"name": project_name}
	update := bson.M{"$set": bson.M{"tags": tags}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *TagsR) DeleteFromProject(ctx context.Context, project_name string, tags ...string) error {
	filter := bson.M{"name": project_name}
	update := bson.M{"$pull": bson.M{"tags": bson.M{"tag_name": bson.M{"$in": tags}}}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}
