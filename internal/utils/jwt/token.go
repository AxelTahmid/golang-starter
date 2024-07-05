package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(data interface{}) (string, error) {

	if data == nil {
		return "", errors.New("data is nil")

	}

	// jwt.ParseECPrivateKeyFromPEM()
	// jwt.ParseECPublicKeyFromPEM()

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
