package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (store *Store) CreateProject(ctx context.Context, project Project) (Project, error) {
	res, err := store.collection.InsertOne(ctx, project)
	if err != nil {
		return Project{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		project.ID = id
	}

	return project, nil
}

func (store *Store) GetProject(ctx context.Context, id primitive.ObjectID) (Project, error) {
	var result Project
	filter := bson.M{"_id": id}
	err := store.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return Project{}, err
	}
	return result, err
}

func (store *Store) GetAllProjects(ctx context.Context, limit, page uint) ([]Project, error) {
	var result []Project
	l := int64(limit)
	skip := int64(page*limit - limit)
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

func (store *Store) DeleteProject(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.collection.DeleteOne(ctx, filter)
	return err
}

type UpdateProject bson.D

func (store *Store) UpdateProject(ctx context.Context, id primitive.ObjectID, arg UpdateProject) (Project, error) {
	update := bson.D{
		{"$set", arg},
	}
	_, err := store.collection.UpdateByID(ctx, id, update)
	if err != nil {
		return Project{}, err
	}
	return store.GetProject(ctx, id)
}
