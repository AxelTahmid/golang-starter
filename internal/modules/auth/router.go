package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/AxelTahmid/golang-starter/db"
)

type AuthHandler struct {
	user UserModel
	v    *validator.Validate
}

func Routes(pg *db.Postgres) chi.Router {
	r := chi.NewRouter()

	authHandler := &AuthHandler{
		user: UserModel{pool: pg.DB},
		v:    validator.New(),
	}

	r.Post("/login", authHandler.login)
	r.Post("/register", authHandler.register)
	// r.Post("/refresh", authHandler.refresh)
	// r.Post("/me", authHandler.me)

	return r
}
