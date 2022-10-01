package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassowrd returns the bcrypt hash of the password
func HashPassowrd(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash passowrd: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
