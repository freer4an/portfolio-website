package helpers

import "net/http"

func SetAdminCookie(w http.ResponseWriter, name, uuid string) {
	c := &http.Cookie{
		Name:     "admin",
		Value:    uuid,
		Secure:   true,
		HttpOnly: true,
		Path:     "/admin",
	}
	http.SetCookie(w, c)
}

func DeleteAdminCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     "admin",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
