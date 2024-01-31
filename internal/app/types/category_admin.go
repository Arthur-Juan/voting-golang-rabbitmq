package types

import "gorm.io/gorm"

type CategoryAdmin struct {
	UserId     uint     `gorm:"primaryKey"`
	CategoryId uint     `gorm:"primaryKey"`
	User       User     `gorm:"foreignKey:UserId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func (c *CategoryAdmin) BeforeCreate(db *gorm.DB) error {
	return nil
}
