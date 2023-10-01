package project

import (
	"context"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/util"
)

const (
	defaultProjectsPage = "projects?page=1"
)

func (api *ProjectAPI) GetAllProjects(w http.ResponseWriter, r *http.Request) {
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

	err = api.temp.ExecuteTemplate(w, "projects.html", projects)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}
