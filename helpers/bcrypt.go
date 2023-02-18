package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword (password string) ([]byte, error) {
	fmt.Println(password)
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword (hash, password string) error {
	fmt.Println(hash)
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}