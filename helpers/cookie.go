package helpers

import "net/http"

func SetCookie(w http.ResponseWriter, uuid string) {
	c := &http.Cookie{
		Name:     "admin",
		Value:    uuid,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)
}

func DeleteCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     "admin",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
