package admin

import (
	"html/template"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/models"
)

var (
	temp *template.Template
)

type Payload struct {
	Projects []models.Project
}

func (api *AdminAPI) Admin(w http.ResponseWriter, r *http.Request) {
	projects, err := api.store.Project.GetAll(r.Context(), 10, 1)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err = temp.ExecuteTemplate(w, "admin.html", projects); err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func init() {
	temp = template.Must(template.ParseGlob("./front/templates/admin/*.html"))
}
