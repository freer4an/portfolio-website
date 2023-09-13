package server

import (
	"errors"
	"fmt"
	"net/http"
)

var session = make(map[string]string)

func (server *Server) Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("admin")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				deleteAdminCookie(w)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			errResponse(w, err, http.StatusForbidden)
			return
		}
		token, ok := session[server.config.AdminName]
		if !ok {
			errResponse(w, fmt.Errorf("Unexpected error: session"), http.StatusInternalServerError)
			return
		}

		if cookie.Name != "admin" || token != cookie.Value {
			deleteAdminCookie(w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
