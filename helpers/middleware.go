package helpers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/freer4an/portfolio-website/internal/repository"
)

var admin_c = os.Getenv("ADMIN_COOKIE")

func MiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(admin_c)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			deleteSession(w, admin_c)
			ErrResponse(w, err, http.StatusForbidden)
			return
		}

		token := repository.GetSessionStr(admin_c)

		if token != cookie.Value {
			deleteSession(w, admin_c)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func deleteSession(w http.ResponseWriter, name string) {
	DeleteAdminCookie(w)
	repository.DeleteSession(name)
}
