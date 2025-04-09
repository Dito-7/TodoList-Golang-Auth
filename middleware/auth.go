package middleware

import (
	"TodoList-Golang-Auth/repository"
	"TodoList-Golang-Auth/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userEmailKey contextKey = "userEmail"

func JWTAuthMiddleware(blacklistRepo repository.BlacklistRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			blacklisted, err := blacklistRepo.IsBlacklisted(r.Context(), tokenString)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			if blacklisted {
				http.Error(w, "Token has been revoked", http.StatusUnauthorized)
				return
			}

			claims, err := utils.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userEmailKey, claims.Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userEmailKey).(string)
	return email, ok
}
