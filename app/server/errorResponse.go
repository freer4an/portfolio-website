package server

import (
	"errors"
	"fmt"
	"net/http"
)

func errResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	if err == nil {
		err = errors.New("business error")
	}
	fmt.Fprintf(w, "%s | %s", http.StatusText(status), err.Error())
}
