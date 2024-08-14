package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/AxelTahmid/golang-starter/api/middlewares"
	"github.com/AxelTahmid/golang-starter/db"
)

type Auth struct {
	db        *db.Postgres
	validator *validator.Validate
}

func Routes(db *db.Postgres) chi.Router {
	r := chi.NewRouter()

	auth := &Auth{
		db:        db,
		validator: validator.New(),
	}

	r.Post("/login", auth.login)
	r.Post("/register", auth.register)

	r.With(middlewares.Authenticated).Get("/me", auth.me)
	r.With(middlewares.AuthenticatedRefreshToken).Post("/refresh", auth.refresh)

	return r
}
