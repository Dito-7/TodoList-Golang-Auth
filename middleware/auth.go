package middleware

import (
	"TodoList-Golang-Auth/models"
	"TodoList-Golang-Auth/utils"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(blacklistRepo models.BlacklistRepository) func(http.Handler) http.Handler {
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

			ctx := utils.SetUserEmailToContext(r.Context(), claims.Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
