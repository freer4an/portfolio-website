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
	r.Get("/", s.welcome)
	r.Get("/projects/{page}", s.projects)
	r.Get("/project/{name}", s.getProjectByName)
	r.Delete("/project/{name}", s.deleteProject)
	r.Patch("/project/{name}", s.updateProject)
	r.Post("/add-project", s.addProject)
	// r_admin := r.Route("/project/{id}", func(r chi.Router) {
	// 	r.Patch("/", s.UpdateProject)
	// 	r.Delete("/", s.DeleteProject)
	// })
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
