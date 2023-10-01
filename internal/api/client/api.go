package client

import (
	"html/template"
	"net/http"

	"github.com/freer4an/portfolio-website/inits/config"
	"github.com/freer4an/portfolio-website/internal/middleware"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/go-chi/chi/v5"
)

type ClientAPI struct {
	store  *repository.Repository
	temp   *template.Template
	config *config.Config
}

func New(store *repository.Repository, temp *template.Template, config *config.Config) *ClientAPI {
	return &ClientAPI{store: store, temp: temp, config: config}
}

func (api *ClientAPI) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", api.Welcome)
	r.Route("/login", func(r chi.Router) {
		r.Get("/", api.Admin_login)
		r.Post("/", api.Login_action)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Get("/admin", api.Admin)
	})
	fs := http.FileServer(http.Dir("./front/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	return r
}
