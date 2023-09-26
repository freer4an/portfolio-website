package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/freer4an/portfolio-website/inits/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	config *config.Config
	router *chi.Mux
	http   *http.Server
}

func NewServer(config *config.Config) *Server {
	server := &Server{config: config}
	return server
}

func (s *Server) InitRoutes(r *chi.Mux) {
	s.router = r
}

func (server *Server) Start(addr string) error {
	switch {
	case server.router == nil:
		err := errors.New("undefined router instance")
		return err
	case addr == "":
		err := errors.New("empty addres")
		return err
	}

	server.http = &http.Server{
		Addr:    addr,
		Handler: server.router,
	}

	log.Info().Msgf("Listening at %s", server.http.Addr)
	return server.http.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down the server")
	return server.http.Shutdown(ctx)
}
