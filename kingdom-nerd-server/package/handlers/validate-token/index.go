package validatetoken

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func validateTokenHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(
		bearerToken[1],
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("My Secret"), nil
		},
	)
	if err != nil {
		http.Error(w, "Error parsing token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		http.Error(w, "Error getting claims: "+errors.New("couldn't parse claims").Error(), http.StatusUnauthorized)
		return
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		http.Error(w, "Error: "+errors.New("jwt is expired").Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "JWT is valid.")
}
