package jwt

import (
	"context"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAccessTokenClaims(token string) (*jwt.RegisteredClaims, error) {
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

	if claims.Issuer != accessTokenIssuer {
		log.Println("invalid token issuer")
		return nil, errTokenInvalid
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

func ParseTokenPair(tokens Tokens) bool {
	accessToken, err := jwt.Parse(tokens.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Printf("error parsing access token -> %v", err)
		return false
	}

	if !accessToken.Valid {
		log.Println("access token expired, parsing refresh token")
	}

	if accessToken.Claims.(jwt.RegisteredClaims).Issuer != accessTokenIssuer {
		log.Println("invalid access token issuer")
		return false
	}

	refreshToken, err := jwt.Parse(tokens.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Printf("error parsing refresh token -> %v", err)
		return false
	}

	if refreshToken.Claims.(jwt.RegisteredClaims).Issuer != refreshTokenIssuer {
		log.Println("invalid refresh token issuer")
		return false
	}

	if !refreshToken.Valid {
		log.Println("refresh token expired, login required")
		return false
	}

	return true
}
