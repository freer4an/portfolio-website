package repository

import (
	"context"

	"github.com/freer4an/portfolio-website/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectI interface {
	Create(ctx context.Context, project models.Project) (primitive.ObjectID, error)
	GetByName(ctx context.Context, name string) (models.Project, error)
	GetAll(ctx context.Context, limit, page int64) ([]models.Project, error)
	Delete(ctx context.Context, name string) (int64, error)
	Update(ctx context.Context, name string, arg interface{}) (interface{}, error)
	// GetProjectByID(ctx context.Context, id primitive.ObjectID) (Project, error)
}

type TagI interface {
	AddToProject(ctx context.Context, project_name string, tags ...models.Tag) error
	DeleteFromProject(ctx context.Context, project_name string, tags ...string) error
}
