package project

import (
	"encoding/json"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/models"
	"github.com/go-chi/chi/v5"
)

func (api *ProjectAPI) addProjectTags(w http.ResponseWriter, r *http.Request) {
	project_name := chi.URLParam(r, "name")

	tags := []models.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}

	err := api.store.Tag.AddToProject(r.Context(), project_name, tags...)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *ProjectAPI) deleteProjectTags(w http.ResponseWriter, r *http.Request) {
	project_name := chi.URLParam(r, "name")

	tags := []string{}
	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}

	err := api.store.Tag.DeleteFromProject(r.Context(), project_name, tags...)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}