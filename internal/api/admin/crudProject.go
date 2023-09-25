package admin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/models"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func (api *AdminAPI) AddProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}
	_, err := api.store.Project.Create(context.TODO(), project)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			helpers.ErrResponse(w, err, http.StatusForbidden)
			return
		}
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *AdminAPI) GetProjectByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	project, err := api.store.Project.GetByName(context.TODO(), name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			helpers.ErrResponse(w, err, http.StatusNotFound)
			return
		}
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, project); err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func (api *AdminAPI) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var req interface{}
	if _, ok := req.(models.Project); !ok {
		helpers.ErrResponse(w, fmt.Errorf("Type assertion error"), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}

	name := chi.URLParam(r, "name")

	_, err := api.store.Project.Update(context.TODO(), name, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			helpers.ErrResponse(w, err, http.StatusNotFound)
			return
		}
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *AdminAPI) DeleteProject(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	count, err := api.store.Project.Delete(context.TODO(), name)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	if count == 0 {
		helpers.ErrResponse(w, nil, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
