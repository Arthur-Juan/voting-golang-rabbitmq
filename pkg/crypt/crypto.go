package crypt

import "golang.org/x/crypto/bcrypt"

func Encrypt(msg string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(msg), 14)
	return string(bytes), err
}

func CheckPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
