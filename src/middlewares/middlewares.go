package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

// Logs
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// User authentication
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.TokenValidate(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
