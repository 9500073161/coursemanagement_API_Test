// middleware/auth.go
package middleware

import (
	"context"
	"coursemanagement/auth"
	"net/http"
	"strings"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Attach user email to context
		ctx := context.WithValue(r.Context(), "userEmail", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
