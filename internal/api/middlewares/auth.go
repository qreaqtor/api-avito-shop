package middlewares

import (
	"context"
	"log/slog"
	"net/http"
)

type tokenManager interface {
	GetContextWithUsername(ctx context.Context, headerValue string) (context.Context, error)
}

type AuthMiddleware struct {
	auth tokenManager
}

func NewAuthMiddleware(auth tokenManager) *AuthMiddleware {
	return &AuthMiddleware{
		auth: auth,
	}
}

func (am *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderValue := r.Header.Get("Authorization")

		ctxWithAuth, err := am.auth.GetContextWithUsername(r.Context(), authHeaderValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			slog.Error(
				"auth middleware",
				"token", authHeaderValue,
				"err", err.Error(),
				"url", r.URL.Path,
				"method", r.Method,
			)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctxWithAuth))
	})
}
