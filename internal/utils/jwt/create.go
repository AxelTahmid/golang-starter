package jwt

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// IssueToken creates a new access and refresh token, should never hold user state, only identifier
func IssueTokenPair(user UserClaims) (Tokens, error) {
	if user.Id == 0 && user.Email == "" && user.Role == "" {
		return Tokens{}, errUserEmpty
	}

	var (
		tokens Tokens
		err    error
	)

	claims := jwt.RegisteredClaims{
		ID:        strconv.Itoa(user.Id),
		Issuer:    accessTokenIssuer,
		Subject:   user.Email,
		Audience:  []string{user.Role},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTime)),
	}

	tokens.AccessToken, err = newToken(claims)
	if err != nil {
		log.Printf("error creating access token -> %v", err)
		return Tokens{}, errTokenCreate
	}

	// longer for refresh token
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(refreshTime))
	claims.Issuer = refreshTokenIssuer

	tokens.RefreshToken, err = newToken(claims)
	if err != nil {
		log.Printf("error creating refresh token -> %v", err)
		return Tokens{}, errTokenCreate
	}

	return tokens, nil
}

func ReIssueAccessToken(user UserClaims) (string, error) {
	if user.Id == 0 && user.Email == "" && user.Role == "" {
		return "", errUserEmpty
	}

	var (
		accessToken string
		err         error
	)

	claims := jwt.RegisteredClaims{
		ID:        strconv.Itoa(user.Id),
		Issuer:    accessTokenIssuer,
		Subject:   user.Email,
		Audience:  []string{user.Role},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTime)),
	}

	accessToken, err = newToken(claims)
	if err != nil {
		log.Printf("error re-issuing access token -> %v", err)
		return "", errTokenCreate
	}

	return accessToken, nil
}
