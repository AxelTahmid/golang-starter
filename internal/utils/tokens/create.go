package tokens

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateToken creates a new access and refresh token, should never hold user state, only identifier
func CreateToken(user UserClaims) (Tokens, error) {
	if user.Id == 0 && user.Email == "" && user.Role == "" {
		return Tokens{}, errors.New("user cannot be empty")
	}

	var (
		tokens Tokens
		err    error
	)

	claims := jwt.RegisteredClaims{
		ID:        strconv.Itoa(user.Id),
		Issuer:    jwtIssuer,
		Subject:   user.Email,
		Audience:  []string{user.Role},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
	}

	tokens.AccessToken, err = newToken(claims)
	if err != nil {
		return Tokens{}, errors.New("error creating access token")
	}

	// longer for refresh token
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 48))

	tokens.RefreshToken, err = newToken(claims)
	if err != nil {
		return Tokens{}, errors.New("error creating refresh token")
	}

	return tokens, nil
}
