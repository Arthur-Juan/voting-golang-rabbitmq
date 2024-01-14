package types

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string
	Description string
	Creator     *User
	CreatorId   uint
	Admins      []*User `gorm:"many2many:admin_category;"`
	Candidates  []*User `gorm:"many2many:canditate_category;"`
}
