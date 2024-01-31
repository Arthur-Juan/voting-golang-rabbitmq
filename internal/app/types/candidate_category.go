package types

import "gorm.io/gorm"

type CandidateCategory struct {
	UserId     uint `gorm:"primaryKey"`
	CategoryId uint `gorm:"primaryKey"`
	Status     Status
	User       User     `gorm:"foreignKey:UserId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func (c *CandidateCategory) BeforeCreate(db *gorm.DB) error {
	return nil
}

type Status uint8

const (
	Pending Status = iota
	Approved
	Rejected
)

type CandidateCategoryOutput struct {
	User     *User           `json:"user"`
	Category *CategoryOutput `json:"category"`
}
