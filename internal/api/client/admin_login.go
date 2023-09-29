package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/repository"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *ClientAPI) Login_action(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusBadRequest)
		return
	}

	if req.Username != viper.GetString("") || req.Password != viper.GetString("") {
		helpers.ErrResponse(w, fmt.Errorf("Failed to confirm data"), http.StatusUnauthorized)
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = repository.AddSession(req.Username, uuid)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	helpers.SetCookie(w, uuid.String())
	w.WriteHeader(http.StatusOK)
}

func (api *ClientAPI) Admin_login(w http.ResponseWriter, r *http.Request) {
	err := api.temp.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		helpers.ErrResponse(w, err, http.StatusInternalServerError)
		return
	}
	return
}

func (api *ClientAPI) Admin_logout(w http.ResponseWriter, r *http.Request) {
	helpers.DeleteCookie(w)
}
