package types

import "gorm.io/gorm"

type CategoryAdmin struct {
	UserId     uint `gorm:"primaryKey"`
	CategoryId uint `gorm:"primaryKey"`
}

func (c *CategoryAdmin) BeforeCreate(db *gorm.DB) error {
	err := db.SetupJoinTable(&User{}, "Categories", &CategoryAdmin{})
	return err
}
