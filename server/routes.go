package server

import (
	"net/http"

	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *Server) init_router() {
	r := chi.NewRouter()
	r.Use(util.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://", "https://"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	s.router = r
}

func (s *Server) BuildRoutes(pattern string, handler http.Handler) {
	s.router.Mount(pattern, handler)
}
