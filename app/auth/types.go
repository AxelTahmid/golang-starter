package auth

import (
	"github.com/AxelTahmid/tinker/db"
	"github.com/AxelTahmid/tinker/pkg/jwt"
)

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterRequest struct {
		db.InsertUser
	}

	LoginResponse struct {
		jwt.Tokens
	}

	RefreshRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RefreshResponse struct {
		AccessToken string `json:"access_token"`
	}
)
