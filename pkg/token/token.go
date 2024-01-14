package token

import (
	"time"

	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	ID    uint64
	Email string
	Name  string
	Exp   int64
}

func GenerateToken(user *types.User) (string, error) {

	claims := &Claims{
		ID:    uint64(user.ID),
		Email: user.Email,
		Name:  user.GetFullName(),
		Exp:   time.Now().Add(time.Hour * 8).Unix(),
	}

	tokenHandler := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenHandler.SignedString([]byte(config.GetKey()))
	if err != nil {
		return "", err
	}

	return token, nil

}
