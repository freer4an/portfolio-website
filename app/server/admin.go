package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

const (
	secret = "slmdasdmad;ladl;afnlka"
)

var admin_temp *template.Template

func (server *Server) admin(w http.ResponseWriter, r *http.Request) {
	projects, err := server.store.GetAllProjects(server.ctx, 10, 1)
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err = admin_temp.Execute(w, projects); err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (server *Server) admin_login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errResponse(w, err, http.StatusBadRequest)
		return
	}

	if req.Username != server.config.AdminName || req.Password != server.config.AdminPass {
		errResponse(w, fmt.Errorf("Failed to confirm data"), http.StatusUnauthorized)
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		errResponse(w, err, http.StatusInternalServerError)
		return
	}

	session[req.Username] = uuid.String()

	cookie := genereateCookie("admin", uuid.String())
	http.SetCookie(w, cookie)
}

func genereateCookie(name, uuid string) *http.Cookie {
	return &http.Cookie{
		Name:     "admin",
		Value:    uuid,
		Secure:   true,
		HttpOnly: true,
		Path:     "/admin",
	}
}

func deleteAdminCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     "admin",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	for k := range session {
		delete(session, k)
	}
}

func init() {
	admin_temp = template.Must(template.ParseFiles("./front/templates/admin/admin.html"))
}
