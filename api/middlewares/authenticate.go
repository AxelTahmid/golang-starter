package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/AxelTahmid/golang-starter/pkg/jwt"
	"github.com/AxelTahmid/golang-starter/pkg/message"
	"github.com/AxelTahmid/golang-starter/pkg/respond"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reply := respond.Write(w)

		// we expect the "Authorization" header to be present in the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			reply.Status(http.StatusUnauthorized).WithErr(message.ErrUnauthorized)
			return
		}

		// we expect the "Authorization" header to be "BEARER {TOKEN}"
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			reply.Status(http.StatusUnauthorized).WithErr(message.ErrBadTokenFormat)
			return
		}

		// parse the JWT string with claims and store the result in `claims`.
		claims, err := jwt.ParseAccessTokenClaims(authHeaderParts[1])

		if err != nil {
			reply.Status(http.StatusBadRequest).WithErr(err)
			return
		}

		// add parsed token data to the request context
		r = r.WithContext(context.WithValue(r.Context(), jwt.AuthReqCtxKey, claims))

		next.ServeHTTP(w, r)
	})
}

func AuthenticateAdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reply := respond.Write(w)

		// get the claims from the request context
		claims, ok := jwt.ParseClaimsCtx(r.Context())
		if !ok {
			reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
			return
		}

		// check if the role is "admin"
		role := claims.Audience
		if role[0] != "admin" {
			reply.Status(http.StatusUnauthorized).WithErr(message.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticatedRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reply := respond.Write(w)

		// we expect the "Authorization" header to be present in the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			reply.Status(http.StatusUnauthorized).WithErr(message.ErrUnauthorized)
			return
		}

		// we expect the "Authorization" header to be "BEARER {TOKEN}"
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			reply.Status(http.StatusUnauthorized).WithErr(message.ErrBadTokenFormat)
			return
		}

		// parse the JWT string with claims and store the result in `claims`.
		claims, err := jwt.ParseRefreshTokenClaims(authHeaderParts[1])

		if err != nil {
			reply.Status(http.StatusBadRequest).WithErr(err)
			return
		}

		// add parsed token data to the request context
		r = r.WithContext(context.WithValue(r.Context(), jwt.AuthReqCtxKey, claims))

		next.ServeHTTP(w, r)
	})
}
