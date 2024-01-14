package authservice

import (
	"errors"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/crypt"
	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/token"
)

func (s *AuthService) Login(input *types.LoginInput) (string, error) {

	var user *types.User

	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", errors.New("wrong credentials")
	}

	err := crypt.CheckPassword(input.Password, user.Password)
	if err != nil {
		return "", errors.New("Wrong credentials")
	}

	token, err := token.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
