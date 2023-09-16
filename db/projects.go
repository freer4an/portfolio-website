package db

import (
	"context"

	"github.com/freer4an/portfolio-website/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (store *Store) CreateProject(ctx context.Context, project models.Project) (primitive.ObjectID, error) {
	res, err := store.collection.InsertOne(ctx, project)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (store *Store) GetProject(ctx context.Context, name string) (models.Project, error) {
	var result models.Project
	filter := bson.M{"name": name}
	err := store.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.Project{}, err
	}
	return result, nil
}

func (store *Store) GetAllProjects(ctx context.Context, limit, page int64) ([]models.Project, error) {
	var result []models.Project
	l := limit
	skip := page*limit - limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	filter := bson.D{}

	cursor, err := store.collection.Find(ctx, filter, &fOpt)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (store *Store) DeleteProject(ctx context.Context, name string) (int64, error) {
	filter := bson.M{"name": name}
	res, err := store.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func (store *Store) UpdateProject(ctx context.Context, name string, arg interface{}) (interface{}, error) {
	filter := bson.M{"name": name}
	update := bson.D{
		{Key: "$set", Value: arg},
	}
	res, err := store.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return res.UpsertedID, nil
}

// func (store *Store) GetProjectByID(ctx context.Context, id primitive.ObjectID) (models.Project, error) {
// 	var result models.Project
// 	filter := bson.M{"_id": id}
// 	err := store.collection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return models.Project{}, err
// 	}
// 	return result, nil
// }
