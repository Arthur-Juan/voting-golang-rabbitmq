package types

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	Creator     *User
	CreatorId   uint
	End         time.Time
	Admins      []*User `gorm:"many2many:admin_category;"`
	Candidates  []*User `gorm:"many2many:canditate_category;"`
}

func NewCategory(name, description string, end time.Time, creator *User) (*Category, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if description == "" {
		return nil, errors.New("description is required")
	}

	if end.IsZero() {
		return nil, errors.New("end date is required")
	}

	now := time.Now()

	if end.After(now) {
		return nil, errors.New("end date cannot be minor then now")
	}

	if creator == nil {
		return nil, errors.New("creator is required")
	}
	return &Category{
		Name:        name,
		Description: description,
		End:         end,
		Creator:     creator,
	}, nil
}

type CreateCategoryInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	End         time.Time `json:"end"`
}
