package auth

import (
	"github.com/go-chi/chi/v5"

	"github.com/AxelTahmid/golang-starter/db"
)

func Routes(pg *db.Postgres) chi.Router {
	r := chi.NewRouter()

	authHandler := AuthHandler{
		postgres: pg,
	}

	r.Post("/login", authHandler.login)
	r.Post("/register", authHandler.register)
	// r.Post("/refresh", authHandler.refresh)
	// r.Post("/me", authHandler.me)

	return r
}
