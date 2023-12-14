package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return passwordHash, nil
}

func Compare(password []byte, hashPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, password)
	return err == nil
}
