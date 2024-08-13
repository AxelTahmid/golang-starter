package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/AxelTahmid/golang-starter/api/middlewares"
	"github.com/AxelTahmid/golang-starter/db"
)

type Auth struct {
	user      UserModel
	validator *validator.Validate
}

func Routes(pg *db.Postgres) chi.Router {
	r := chi.NewRouter()

	auth := &Auth{
		user:      UserModel{pool: pg.Conn()},
		validator: validator.New(),
	}

	r.Post("/login", auth.login)
	r.Post("/register", auth.register)

	r.With(middlewares.Authenticated).Get("/me", auth.me)
	r.With(middlewares.AuthenticatedRefreshToken).Post("/refresh", auth.refresh)

	return r
}
