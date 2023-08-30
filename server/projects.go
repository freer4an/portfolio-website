package server

import (
	"net/http"
	"text/template"

	"github.com/freer4an/portfolio-website/util"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (server *Server) AddProject(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) GetProjects(w http.ResponseWriter, r *http.Request) {
	pageParam := chi.URLParam(r, "page")
	page, err := util.UrlParamToInt(pageParam)
	if err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	projects, err := server.store.GetAllProjects(server.ctx, 5, page)
	t, err := template.ParseFiles("./front/templates/projects.page.html")
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	t.Execute(w, projects)
}

func (server *Server) GetProjectByIDs(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	// server.store.GetProject()

}

func (server *Server) UpdateProject(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) DeleteProject(w http.ResponseWriter, r *http.Request) {

}
