package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const saltRounds int = 13

// HashPassword using crypto
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	return string(bytes), err
}

// CheckPasswordHash using crypto
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(password, hash)
	return err == nil
}
