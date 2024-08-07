package auth

import (
	"github.com/AxelTahmid/golang-starter/pkg/jwt"
)

type (
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
		jwt.Tokens
	}

	RefreshRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RefreshResponse struct {
		AccessToken string `json:"access_token"`
	}
)
