package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ...

// Middleware to authenticate users using JWT
func AuthenticateMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	secretKey := "3958d380-517b-44a5-9429-5f53cddacae5"
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Expose-Headers", "*")
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing auth token", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})

			if err != nil {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims)
			} else {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}
