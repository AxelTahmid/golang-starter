package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/AxelTahmid/golang-starter/api/middlewares"
	"github.com/AxelTahmid/golang-starter/db"
)

type AuthHandler struct {
	user UserModel
	v    *validator.Validate
}

func Routes(pg *db.Postgres) chi.Router {
	r := chi.NewRouter()

	authHandler := &AuthHandler{
		user: UserModel{pool: pg.Conn()},
		v:    validator.New(),
	}

	r.Post("/login", authHandler.login)
	r.Post("/register", authHandler.register)

	r.With(middlewares.Authenticated).Get("/me", authHandler.me)
	r.With(middlewares.AuthenticatedRefreshToken).Post("/refresh", authHandler.refresh)

	return r
}
