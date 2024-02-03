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

type GrantAdminAccessInput struct {
	CategoryId int `json:"category_id"`
	TargetId   int `json:"target_id"`
}
