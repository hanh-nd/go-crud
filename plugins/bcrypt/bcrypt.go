package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func Compare(password string, hashedPassword string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if result != nil {
		return false
	}
	return true
}
