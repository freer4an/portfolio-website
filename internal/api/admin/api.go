package admin

import (
	"github.com/freer4an/portfolio-website/internal/middleware"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/go-chi/chi/v5"
)

type AdminAPI struct {
	store *repository.Repository
}

func New(store *repository.Repository) *AdminAPI {
	return &AdminAPI{store: store}
}

// admin routes
func (api *AdminAPI) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/login", func(r chi.Router) {
		r.Get("/", api.Admin_login)
		r.Post("/", api.Login_action)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Get("/", api.Admin)
		r.Route("/projects", func(r chi.Router) {
			r.Post("/", api.AddProject)
			r.Route("/{name}", func(r chi.Router) {
				r.Delete("/", api.DeleteProject)
				r.Patch("/", api.UpdateProject)
				r.Get("/", api.GetProjectByName)
			})
		})
	})
	return r
}
