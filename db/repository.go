package db

import (
	"context"

	"github.com/freer4an/portfolio-website/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	ProjectI
	TagI
}

type ProjectI interface {
	CreateProject(ctx context.Context, project models.Project) (primitive.ObjectID, error)
	GetProject(ctx context.Context, name string) (models.Project, error)
	GetAllProjects(ctx context.Context, limit, page int64) ([]models.Project, error)
	DeleteProject(ctx context.Context, name string) (int64, error)
	UpdateProject(ctx context.Context, name string, arg interface{}) (interface{}, error)
	// GetProjectByID(ctx context.Context, id primitive.ObjectID) (Project, error)
}

type TagI interface {
	AddProjectTags(ctx context.Context, project_name string, tags ...models.Tag) error
	DeleteProjectTags(ctx context.Context, project_name string, tags ...string) error
}
