package server

import (
	"encoding/json"
	"net/http"

	"github.com/freer4an/portfolio-website/db"
	"github.com/go-chi/chi/v5"
)

func (server *Server) addProjectTags(w http.ResponseWriter, r *http.Request) {
	project_name := chi.URLParam(r, "name")

	tags := []db.Tag{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags); err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	err := server.store.AddProjectTags(server.ctx, project_name, tags...)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) deleteProjectTags(w http.ResponseWriter, r *http.Request) {
	project_name := chi.URLParam(r, "name")

	tags := []string{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags); err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	err := server.store.DeleteProjectTags(server.ctx, project_name, tags...)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
