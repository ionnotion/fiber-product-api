package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secret = "rahasia"

func GenerateToken (UserId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["UserId"] = UserId
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken (c *fiber.Ctx) error {
	authorization :=  c.Get("Authorization", "nil")

	if authorization == "nil" {
		return c.Context().Err()
	}

	tokenString := strings.Split(authorization, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v",token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		UserId, err := strconv.ParseInt(fmt.Sprintf("%.0f",claims["UserId"]),10,32)
		
		if err != nil {
			return err
		}

		c.Locals("id",UserId)
	}

	// fmt.Println(token)

	return nil
}