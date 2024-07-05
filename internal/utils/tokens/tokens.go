package tokens

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtIssuer           = "auth-service"
	refreshTokenSubject = "refresh-token"
	accessTokenSubject  = "access-token"
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
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
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

// func parseAccessToken(accessToken string) *UserClaims {
// 	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})

// 	return parsedAccessToken.Claims.(*UserClaims)
// }

// func parseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
// 	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})

// 	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
// }

// func VerifyTokens(tokens Tokens) error {
// 	var err error

// 	userClaims := parseAccessToken(tokens.AccessToken)
// 	refreshClaims := parseRefreshToken(tokens.RefreshToken)

// 	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 	// 	return secretKey, nil
// 	// })

// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// if !token.Valid {
// 	// 	return fmt.Errorf("invalid token")
// 	// }

// 	// refresh token is expired
// 	if refreshClaims.Valid() != nil {
// 		tokens.RefreshToken, err = newRefreshToken(*refreshClaims)
// 		if err != nil {
// 			log.Fatal("error creating refresh token")
// 		}
// 	}

// 	// access token is expired
// 	if userClaims.RegisteredClaims.Valid() != nil && refreshClaims.Valid() == nil {
// 		tokens.AccessToken, err = newToken(*userClaims)
// 		if err != nil {
// 			log.Fatal("error creating access token")
// 		}
// 	}

// 	return nil
// }
