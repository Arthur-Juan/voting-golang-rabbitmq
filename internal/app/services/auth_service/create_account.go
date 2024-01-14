package authservice

import (
	"errors"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/crypt"
	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/token"
	"gorm.io/gorm"
)

func (s *AuthService) CreateAccount(input *types.CreateAccountInput) (string, error) {

	var user *types.User
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("email already in use")
		}
	}

	if input.Password != input.ConfirmPassword {
		return "", errors.New("passwords are not equal")
	}

	encrypted_pass, err := crypt.Encrypt(input.Password)
	if err != nil {
		return "", err
	}

	user, err = types.NewUser(input.FirstName, input.LastName, input.Email, encrypted_pass)

	if err != nil {
		return "", err
	}

	s.db.Create(&user)
	token, err := token.GenerateToken(user)
	return token, nil
}
