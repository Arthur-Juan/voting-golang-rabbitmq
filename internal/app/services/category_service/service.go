package categoryservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		db: config.GetDb(),
	}
}
