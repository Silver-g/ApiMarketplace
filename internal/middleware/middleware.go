package middleware

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/security"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			boundary.WriteResponseErr(w, http.StatusUnauthorized, boundary.ErrorResponse{
				ErrorCode: "Unauthorized",
				Message:   "Missing Authorization header",
			})
			return
		}

		tokenStr := strings.Fields(authHeader)
		if len(tokenStr) != 2 || tokenStr[0] != "Bearer" {
			boundary.WriteResponseErr(w, http.StatusUnauthorized, boundary.ErrorResponse{
				ErrorCode: "Unauthorized",
				Message:   "Invalid Authorization header format",
			})
			return
		}

		userID, err := security.ParseJwt(tokenStr[1])
		if err != nil {
			boundary.WriteResponseErr(w, http.StatusUnauthorized, boundary.ErrorResponse{
				ErrorCode: "Unauthorized",
				Message:   "Invalid token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func OptionalAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		var userID int

		if authHeader != "" {
			tokenStr := strings.Fields(authHeader)
			if len(tokenStr) == 2 && tokenStr[0] == "Bearer" {
				id, err := security.ParseJwt(tokenStr[1])
				if err == nil {
					userID = id
				}
			}
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
