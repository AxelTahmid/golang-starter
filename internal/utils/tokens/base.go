package tokens

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtIssuer = "auth-service"
)

type UserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// jwt.ParseECPrivateKeyFromPEM()
// jwt.ParseECPublicKeyFromPEM()
// SigningMethodES256

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var (
	errTokenExpired     = errors.New("token expired")
	errSignatureInvalid = errors.New("invalid token signature")
	errParsingToken     = errors.New("error parsing token")
	errUserEmpty        = errors.New("user cannot be empty")
	errTokenCreate      = errors.New("error creating token")
)

func newToken(claims jwt.RegisteredClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
