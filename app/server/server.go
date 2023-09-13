package server

import (
	"context"
	"net/http"

	"github.com/freer4an/portfolio-website/db"
	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	config util.Config
	router *chi.Mux
	store  *db.Store
	ctx    context.Context
	http   *http.Server
}

func NewServer(ctx context.Context, config util.Config, store *db.Store) *Server {
	server := &Server{ctx: ctx, config: config, store: store}
	server.initRoutes()
	return server
}

func (s *Server) initRoutes() {
	r := chi.NewRouter()
	r.Use(logger)

	// admin routes
	r.Route("/admin", func(r chi.Router) {
		r.Get("/", s.admin)
		r.Post("/login", s.admin_login)
		r.Post("/projects", s.addProject)
		r.Delete("/projects/{name}", s.deleteProject)
		r.Patch("/projects/{name}", s.updateProject)
	})

	// public routes
	r.Route("/", func(r chi.Router) {
		r.Get("/", s.welcome)
		r.Get("/projects", s.projects)
		r.Get("/{name}", s.getProjectByName)
	})

	fs := http.FileServer(http.Dir("./front/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	s.router = r
}

func (server *Server) Start() error {
	server.http = &http.Server{
		Addr:    server.config.HttpAddr,
		Handler: server.router,
	}
	log.Info().Msgf("Listening at %s", server.http.Addr)
	return server.http.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down the server")
	return server.http.Shutdown(ctx)
}
