package client

import (
	"context"
	"html/template"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/util"
)

var temp *template.Template

const (
	defaultProjectsPage = "projects?page=1"
)

func (api *ProjectAPI) Projects(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	page, err := util.UrlParamToInt(pageParam)
	if err != nil {
		http.Redirect(w, r, defaultProjectsPage, http.StatusMovedPermanently)
		return
	}

	projects, err := api.store.Project.GetAll(context.TODO(), 4, page)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = temp.ExecuteTemplate(w, "projects.html", projects)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func init() {
	temp = template.Must(template.ParseGlob("./front/templates/*.html"))
}
