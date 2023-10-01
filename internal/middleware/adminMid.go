package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/freer4an/portfolio-website/helpers"
	"github.com/freer4an/portfolio-website/internal/repository"
)

var admin_c = "admin"

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(admin_c)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			deleteSession(w)
			helpers.ErrResponse(w, err, http.StatusForbidden)
			return
		}

		token, err := repository.GetSessionStr(admin_c)
		if err != nil {
			helpers.ErrResponse(w, err, http.StatusUnauthorized)
			return
		}

		if token != cookie.Value {
			deleteSession(w)
			helpers.ErrResponse(w, fmt.Errorf("invalid cookie"), http.StatusUnauthorized)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func deleteSession(w http.ResponseWriter) {
	helpers.DeleteCookie(w, admin_c)
	repository.DeleteSession(admin_c)
}
