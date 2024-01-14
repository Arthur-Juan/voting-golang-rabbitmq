package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitPg()
	if err != nil {
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetKey() string {
	return "change-me"
}
