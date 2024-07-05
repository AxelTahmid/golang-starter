package tokens

import (
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

func newToken(claims jwt.RegisteredClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
