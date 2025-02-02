package db

import (
	"user-service/internal/domain/users"

	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email, password string) (users.User, bool) {
	u, found, _ := UserAlreadyExists(email)
	if !found {
		return u, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return u, false
	}

	return u, true
}