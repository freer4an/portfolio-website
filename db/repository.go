package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	ProjectI
	TagI
}

type ProjectI interface {
	CreateProject(ctx context.Context, project Project) (primitive.ObjectID, error)
	GetProject(ctx context.Context, name string) (Project, error)
	GetAllProjects(ctx context.Context, limit, page int64) ([]Project, error)
	DeleteProject(ctx context.Context, name string) (int64, error)
	UpdateProject(ctx context.Context, name string, arg interface{}) (interface{}, error)
	// GetProjectByID(ctx context.Context, id primitive.ObjectID) (Project, error)
}

type TagI interface {
	AddProjectTags(ctx context.Context, project_name string, tags ...Tag) error
	DeleteProjectTags(ctx context.Context, project_name string, tags ...string) error
}
