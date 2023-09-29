package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router *chi.Mux
	http   *http.Server
}

func New() *Server {
	server := &Server{}
	server.init_router()
	return server
}

func (server *Server) Start(addr string) error {
	switch {
	case server.router == nil:
		err := errors.New("nil router")
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
