package server

import (
	"net/http"
	"text/template"

	"github.com/freer4an/portfolio-website/util"
)

var temp *template.Template

func (server *Server) welcome(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "welcome.html", nil)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func (server *Server) projects(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	page, err := util.UrlParamToInt(pageParam)
	if err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	projects, err := server.store.GetAllProjects(server.ctx, 5, page)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = temp.ExecuteTemplate(w, "projects.html", projects)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func init() {
	temp = template.Must(template.ParseGlob("./front/templates/*.html"))
}
