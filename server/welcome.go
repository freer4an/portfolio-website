package server

import (
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

func (server *Server) Welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./front/templates/welcome.html")
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	t.Execute(w, nil)
}
