package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const SALT_ROUNDS int = 13;

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), SALT_ROUNDS)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(password, hash)
	return err == nil
}

