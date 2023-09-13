package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/freer4an/portfolio-website/db"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func (server *Server) addProject(w http.ResponseWriter, r *http.Request) {
	var project db.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}
	_, err := server.store.CreateProject(server.ctx, project)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			errResponse(w, err, http.StatusForbidden)
			return
		}
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) getProjectByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	project, err := server.store.GetProject(server.ctx, name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errResponse(w, err, http.StatusNotFound)
			return
		}
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err := temp.ExecuteTemplate(w, "project.html", project); err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func (server *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	var project interface{}
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	name := chi.URLParam(r, "name")

	_, err := server.store.UpdateProject(server.ctx, name, project)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errResponse(w, err, http.StatusNotFound)
			return
		}
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) deleteProject(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	count, err := server.store.DeleteProject(server.ctx, name)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	if count == 0 {
		errResponse(w, nil, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
