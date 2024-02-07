package types

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Voter      *User
	VoterId    uint
	Category   *Category
	CategoryId uint
}

type VoteInput struct {
	TargetId   uint `json:"target_id"`
	CategoryId uint `json:"category_id"`
}
