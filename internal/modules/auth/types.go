package auth

import (
	"time"

	"github.com/AxelTahmid/golang-starter/db"
)

type (
	AuthHandler struct {
		postgres *db.Postgres
	}

	AuthService struct{}

	// exact order as in database
	User struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Verified  bool      `json:"verified"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
