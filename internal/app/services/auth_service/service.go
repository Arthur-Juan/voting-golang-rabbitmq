package authservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		db: config.GetDb(),
	}
}
