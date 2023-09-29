package client

import (
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/models"
)

type Payload struct {
	Projects []models.Project
}

func (api *ClientAPI) Admin(w http.ResponseWriter, r *http.Request) {
	projects, err := api.store.Project.GetAll(r.Context(), 10, 1)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err = api.temp.ExecuteTemplate(w, "admin.html", projects); err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}
