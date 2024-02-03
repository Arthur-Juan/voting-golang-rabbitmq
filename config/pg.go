package config

import (
	"os"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPg() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := connect(dsn, 0)
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

func connect(dsn string, count int) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if count < 3 {
			count++
			connect(dsn, count)
		}
		return nil, err
	}
	return db, err
}
