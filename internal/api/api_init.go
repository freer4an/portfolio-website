package api

import (
	"html/template"

	"github.com/freer4an/portfolio-website/internal/api/client"
	"github.com/freer4an/portfolio-website/internal/api/project"
	"github.com/freer4an/portfolio-website/internal/repository"
)

type API struct {
	Client  *client.ClientAPI
	Project *project.ProjectAPI
}

func New(store *repository.Repository, temp *template.Template) *API {
	clientAPI := client.New(store, temp)
	projectAPI := project.New(store, temp)
	return &API{
		Client:  clientAPI,
		Project: projectAPI,
	}
}
