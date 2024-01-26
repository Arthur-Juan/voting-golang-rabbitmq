package types

import "gorm.io/gorm"

type CandidateCategory struct {
	UserId     uint `gorm:"primaryKey"`
	CategoryId uint `gorm:"primaryKey"`
	Status     Status
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
