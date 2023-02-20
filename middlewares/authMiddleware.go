package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/security"
)

type contextKey string

const TokenClaims contextKey = "TokenClaims"

func MiddlewareAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenClaims, err := security.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), "decoded", tokenClaims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
