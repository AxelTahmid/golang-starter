package jwt

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// IssueToken creates a new access and refresh token, should never hold user state, only identifier
func IssueToken(user UserClaims) (Tokens, error) {
	if user.Id == 0 && user.Email == "" && user.Role == "" {
		return Tokens{}, errUserEmpty
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
		log.Printf("error creating access token -> %v", err)
		return Tokens{}, errTokenCreate
	}

	// longer for refresh token
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 48))

	tokens.RefreshToken, err = newToken(claims)
	if err != nil {
		log.Printf("error creating refresh token -> %v", err)
		return Tokens{}, errTokenCreate
	}

	return tokens, nil
}
