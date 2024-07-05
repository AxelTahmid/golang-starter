package auth

import (
	"github.com/AxelTahmid/golang-starter/internal/utils/tokens"
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
		tokens.Tokens
	}
)
