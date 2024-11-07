package middleware

import (
	"awesomeProject/internal/auth/jwt"
	"context"
	"fmt"
	"net/http"
	"strings"
)

const jwtSecret = "awesomeProject"

const partWithToken = 1

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")[partWithToken]

		claims, err := jwt.Verify(jwtSecret, token)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userId, err := claims.GetSubject()
		if err != nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
