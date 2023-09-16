package db

import (
	"context"

	"github.com/freer4an/portfolio-website/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (store *Store) AddProjectTags(ctx context.Context, project_name string, tags ...models.Tag) error {
	filter := bson.M{"name": project_name}
	update := bson.M{"$set": bson.M{"tags": tags}}
	_, err := store.collection.UpdateOne(ctx, filter, update)
	return err
}

func (store *Store) DeleteProjectTags(ctx context.Context, project_name string, tags ...string) error {
	filter := bson.M{"name": project_name}
	update := bson.M{"$pull": bson.M{"tags": bson.M{"tag_name": bson.M{"$in": tags}}}}
	_, err := store.collection.UpdateOne(ctx, filter, update)
	return err
}
