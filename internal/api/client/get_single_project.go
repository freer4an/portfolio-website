package client

import (
	"errors"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func (api *ProjectAPI) GetProjectByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	project, err := api.store.Project.GetByName(r.Context(), name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			helpers.ErrResponse(w, err, http.StatusNotFound)
			return
		}
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = temp.ExecuteTemplate(w, "project.html", project)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}
