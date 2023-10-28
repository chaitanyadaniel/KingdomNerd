package authutils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(authHeader string) (bool, error) {
	if authHeader == "" {
		return false, errors.New("Authorization header is missing")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return false, errors.New("Invalid Authorization header")
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
		return false, errors.New("Error parsing token: " + err.Error())
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false, errors.New("Error getting claims: " + errors.New("couldn't parse claims").Error())
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return false, errors.New("JWT is expired")
	}

	return true, nil
}
