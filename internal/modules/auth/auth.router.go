package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {

	r := chi.NewRouter()

	r.Post("/login", loginHandler)
	// r.Post("/register", RegisterHandler)
	// r.Post("/refresh", RefreshHandler)

	return r
}
