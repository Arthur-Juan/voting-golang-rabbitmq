package types

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Voter      *User
	VoterId    uint
	Category   *Category
	CategoryId uint
	Target     *User
	TargetId   uint
}
