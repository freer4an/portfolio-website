package server

import (
	"fmt"
	"net/http"
)

func errResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s | %s", http.StatusText(status), err.Error())
}
