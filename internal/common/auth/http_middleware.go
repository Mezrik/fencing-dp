package auth

import (
	"net/http"
	"strings"

	"github.com/Mezrik/fencing-dp/internal/common/errors"
)

type AuthHttpMiddleware struct {
}

func (a AuthHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			errors.Unauthorised("empty-bearer-token", nil, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a AuthHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}
