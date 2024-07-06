package tokens

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(token string) (*jwt.RegisteredClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("invalid token signature -> %v", err)
			return nil, errSignatureInvalid
		}
		log.Printf("error parsing token -> %v", err)
		return nil, errParsingToken
	}

	if !parsedToken.Valid {
		log.Printf("expired token -> %v", err)
		return nil, errTokenExpired
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)

	if !ok {
		log.Println("error parsing claims")
		return nil, errParsingToken
	}

	return claims, nil
}

// func RefreshTokens(tokens Tokens) error {
// 	var err error

// 	accessToken, accErr := ParseToken(tokens.AccessToken)
// 	refreshClaims, refErr := ParseToken(tokens.RefreshToken)

// 	// refresh token is expired
// 	if refErr != nil {
// 		refreshClaims, err = newToken(*refreshClaims)
// 		if err != nil {
// 			log.Fatal("error creating refresh token")
// 		}
// 	}

// 	// access token is expired
// 	if accErr != nil && refErr == nil {
// 		accessToken, err = newToken(*accessClaims)
// 		if err != nil {
// 			log.Fatal("error creating access token")
// 		}
// 	}

// 	return nil
// }
