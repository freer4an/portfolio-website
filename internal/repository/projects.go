package repository

import (
	"context"

	"github.com/freer4an/portfolio-website/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectsR struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProjectR(db *mongo.Database, collection *mongo.Collection) *ProjectsR {
	return &ProjectsR{
		db:         db,
		collection: collection,
	}
}

func (repo *ProjectsR) Create(ctx context.Context, project models.Project) (primitive.ObjectID, error) {
	res, err := repo.collection.InsertOne(ctx, project)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo *ProjectsR) GetByName(ctx context.Context, name string) (models.Project, error) {
	var result models.Project
	filter := bson.M{"name": name}
	err := repo.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.Project{}, err
	}
	return result, nil
}

func (repo *ProjectsR) GetAll(ctx context.Context, limit, page int64) ([]models.Project, error) {
	var result []models.Project
	l := limit
	skip := page*limit - limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	filter := bson.D{}

	cursor, err := repo.collection.Find(ctx, filter, &fOpt)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (repo *ProjectsR) Delete(ctx context.Context, name string) (int64, error) {
	filter := bson.M{"name": name}
	res, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func (repo *ProjectsR) Update(ctx context.Context, name string, arg interface{}) (interface{}, error) {
	filter := bson.M{"name": name}
	update := bson.D{
		{Key: "$set", Value: arg},
	}
	res, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return res.UpsertedID, nil
}

// func (repo *ProjectsR) GetByID(ctx context.Context, id primitive.ObjectID) (models.Project, error) {
// 	var result models.Project
// 	filter := bson.M{"_id": id}
// 	err := repo.collection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		return models.Project{}, err
// 	}
// 	return result, nil
// }
