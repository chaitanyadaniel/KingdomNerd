package authutils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Function to create a new JWT
func CreateToken(username, firstName, lastName string) (string, error) {

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["first_name"] = firstName
	atClaims["last_name"] = lastName
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("3958d380-517b-44a5-9429-5f53cddacae5"))
	if err != nil {
		return "", err
	}
	return token, nil
}
