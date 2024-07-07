package jwt

import (
	"context"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(token string) (*jwt.RegisteredClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Printf("ParseWithClaims errored -> %v", err)
		return nil, err
	}

	if !parsedToken.Valid {
		log.Printf("invalid token -> %v", err)
		return nil, errTokenInvalid
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)

	if !ok {
		log.Println("error parsing claims")
		return nil, errParsingClaims
	}

	return claims, nil
}

func ParseClaimsCtx(ctx context.Context) (*jwt.RegisteredClaims, bool) {
	userClaim, ok := ctx.Value(AuthReqCtxKey).(*jwt.RegisteredClaims)
	if !ok {
		log.Println("error parsing claims")
		return nil, false
	}

	return userClaim, true
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
