package security

import (
	"ApiMarketplace/internal/consts"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordHashing = errors.New(consts.ErrPasswordHashingMsg)

// для регистрации
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrPasswordHashing
	}
	return string(hash), nil
}

// для логина
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
