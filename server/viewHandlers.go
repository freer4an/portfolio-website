package server

import (
	"net/http"
	"text/template"

	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
)

var temp *template.Template

func (server *Server) welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./front/templates/welcome.html")
	if err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}
	t.Execute(w, nil)
}

func (server *Server) projects(w http.ResponseWriter, r *http.Request) {
	pageParam := chi.URLParam(r, "page")
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
