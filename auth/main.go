package auth

import (
	"github.com/go-chi/chi"
	"net/http"
)

// Handler for all auth purposes
func Handler() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", Login)
	return r
}
