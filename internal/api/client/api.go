package client

import (
	"net/http"

	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/go-chi/chi/v5"
)

type ProjectAPI struct {
	store *repository.Repository
}

func New(store *repository.Repository) *ProjectAPI {
	return &ProjectAPI{store: store}
}

// project routes
func (api *ProjectAPI) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", api.Welcome)
		r.Get("/projects", api.Projects)
		r.Get("/projects/{name}", api.GetProjectByName)
	})
	fs := http.FileServer(http.Dir("./front/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	return r
}
