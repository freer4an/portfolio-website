package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Err CustomError `json:"error"`
}

type CustomError struct {
	StatusCode int    `json:"error_code"`
	StatusText string `json:"error_text"`
	Msg        string `json:"error_msg"`
}

func errResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		err = errors.New("usecase error")
	}
	cerr := CustomError{
		StatusCode: status,
		StatusText: http.StatusText(status),
		Msg:        err.Error(),
	}
	resp := ErrorResponse{cerr}
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
	}
}
