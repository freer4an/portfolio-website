package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CustomError struct {
	StatusCode int    `json:"error_code"`
	StatusText string `json:"error_text"`
	Msg        string `json:"error_msg"`
}

func ErrResponse(w http.ResponseWriter, err error, status int) {
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

	if err = json.NewEncoder(w).Encode(&cerr); err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
	}
}
