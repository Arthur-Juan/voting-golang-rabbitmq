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
	Winners     uint
	Admins      []*User `gorm:"many2many:admin_categories;"`
	Candidates  []*User `gorm:"many2many:canditate_categories;"`
}

func NewCategory(name, description string, end time.Time, creator *User, winners int) (*Category, error) {

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

	if end.Before(now) {
		return nil, errors.New("end date cannot be minor then now")
	}

	if winners <= 0 {
		return nil, errors.New("winner cannot be minor then 0")
	}

	if creator == nil {
		return nil, errors.New("creator is required")
	}
	return &Category{
		Name:        name,
		Description: description,
		End:         end,
		Creator:     creator,
		Winners:     uint(winners),
	}, nil
}

type CreateCategoryInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	End         time.Time `json:"end"`
	Winners     int       `json:"winners"`
}

type CategoryOutput struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	End         time.Time `json:"end"`
	Winners     int       `json:"total_winners"`
	Status      string    `json:"status"`
}

type InviteToCategoryInput struct {
	CategoryId uint `json:"category_id"`
	InviteType uint `json:"invite_type"`
}
