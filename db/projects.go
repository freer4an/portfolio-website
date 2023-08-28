package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *Store) CreateProject(ctx context.Context, project Project) (Project, error) {
	coll := store.db.Collection(projects)
	res, err := coll.InsertOne(ctx, project)
	if err != nil {
		return Project{}, fmt.Errorf("Error creating the project: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		project.ID = id
	}

	return project, nil
}

func (store *Store) GetProject(ctx context.Context, id primitive.ObjectID) (Project, error) {
	var result Project
	coll := store.db.Collection(projects)
	filter := bson.M{"_id": id}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return Project{}, fmt.Errorf("Error getting the project: %w", err)
	}
	return result, err
}

func (store *Store) DeleteProject(ctx context.Context, id primitive.ObjectID) error {
	coll := store.db.Collection(projects)
	filter := bson.M{"_id": id}
	_, err := coll.DeleteOne(ctx, filter)
	return err
}

type updateProject struct {
	Name        string `bson:"name"`
	Description string `bson:"description, omitempty"`
	Url         string `bson:"url, omitempty"`
	IsFinished  bool   `bson:"finished, omitempty"`
}

func (store *Store) UpdateProject(ctx context.Context, id primitive.ObjectID, arg updateProject) (Project, error) {
	coll := store.db.Collection(projects)
	update := bson.D{
		{"$set",
			bson.D{
				{"name", arg.Name},
				{"description", arg.Description},
				{"url", arg.Url},
				{"finished", arg.IsFinished},
			}},
	}
	_, err := coll.UpdateByID(ctx, id, update)
	if err != nil {
		return Project{}, fmt.Errorf("Error updating the project %w", err)
	}
	return store.GetProject(ctx, id)
}
