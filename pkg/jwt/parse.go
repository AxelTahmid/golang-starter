package jwt

import (
	"context"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAccessTokenClaims(token string) (*jwt.RegisteredClaims, error) {
	return parseClaims(token, accessTokenIssuer)
}

func ParseRefreshTokenClaims(token string) (*jwt.RegisteredClaims, error) {
	return parseClaims(token, refreshTokenIssuer)
}

func ParseClaimsCtx(ctx context.Context) (*jwt.RegisteredClaims, bool) {
	userClaim, ok := ctx.Value(AuthReqCtxKey).(*jwt.RegisteredClaims)
	if !ok {
		log.Println("error parsing claims")
		return nil, false
	}

	return userClaim, true
}

func parseClaims(token string, issuer string) (*jwt.RegisteredClaims, error) {
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
		return nil, errParsingClaims
	}

	if claims.Issuer != issuer {
		return nil, errTokenIssuer
	}

	return claims, nil
}
