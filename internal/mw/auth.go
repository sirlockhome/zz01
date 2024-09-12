package mw

import (
	"context"
	"foxomni/pkg/jwt"
	"net/http"
	"strings"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func (mw *Middleware) ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		claims := &jwt.JwtCustomClaims{}

		token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
			return []byte(mw.conf.Auth.AccessKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
