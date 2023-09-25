package admin

import (
	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
)

type AdminAPI struct {
	config *util.Config
	store  *repository.Repository
}

func New(config *util.Config, store *repository.Repository) *AdminAPI {
	return &AdminAPI{config: config, store: store}
}

// admin routes
func (api *AdminAPI) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/login", func(r chi.Router) {
		r.Get("/", api.Admin_login)
		r.Post("/", api.Login_action)
	})
	r.Group(func(r chi.Router) {
		r.Use(helpers.MiddlewareAdmin)
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
