package project

import (
	"html/template"

	"github.com/freer4an/portfolio-website/internal/middleware"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/go-chi/chi/v5"
)

type ProjectAPI struct {
	store *repository.Repository
	temp  *template.Template
}

func New(store *repository.Repository, temp *template.Template) *ProjectAPI {
	return &ProjectAPI{store: store, temp: temp}
}

func (api *ProjectAPI) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", api.GetAllProjects)
	r.Get("/{name}", api.GetProjectByName)
	r.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Post("/", api.AddProject)
		r.Route("/{name}", func(r chi.Router) {
			r.Patch("/", api.UpdateProject)
			r.Delete("/", api.DeleteProject)
		})
	})
	return r
}
