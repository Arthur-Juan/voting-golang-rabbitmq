package config

import (
	"os"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPg() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.SetupJoinTable(&types.User{}, "Categories", &types.CandidateCategory{})

	err = db.AutoMigrate(&types.User{}, &types.Category{}, &types.Vote{}, &types.CandidateCategory{}, &types.CategoryAdmin{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
