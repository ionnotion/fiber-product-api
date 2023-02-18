package helpers

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = "rahasia"

func GenerateToken (UserId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["UserId"] = UserId
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(secret))
}