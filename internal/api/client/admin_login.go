package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/google/uuid"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const session_admin = "admin"

func (api *ClientAPI) Login_action(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}

	if req.Username != api.config.Admin.Login || req.Password != api.config.Admin.Password {
		helpers.ErrResponse(w, fmt.Errorf("Failed to confirm data"), http.StatusUnauthorized)
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	repository.AddSession(session_admin, uuid)

	helpers.SetCookie(w, session_admin, uuid.String())
	w.WriteHeader(http.StatusOK)
}

func (api *ClientAPI) Admin_login(w http.ResponseWriter, r *http.Request) {
	if err := api.temp.ExecuteTemplate(w, "login.html", nil); err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func (api *ClientAPI) Admin_logout(w http.ResponseWriter, r *http.Request) {
	helpers.DeleteCookie(w, session_admin)
}
