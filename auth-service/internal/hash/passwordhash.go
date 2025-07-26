package hash

import (
	"net/http"
	"vitals-guard/common/errs"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errs.New("password must not be empty", "PASSWORD_EMPTY", http.StatusBadRequest)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword, password string) error {
	if hashedPassword == "" || password == "" {
		return errs.New("password must not be empty", "PASSWORD_EMPTY", http.StatusBadRequest)

	}

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
