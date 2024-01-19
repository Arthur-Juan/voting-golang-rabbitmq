package voteservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
)

func (s *VoteService) CreateCategory(input *types.CreateCategoryInput, user_id uint) error {
	var user *types.User

	if err := s.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return err
	}

	category, err := types.NewCategory(input.Name, input.Description, input.End, user)
	if err != nil {
		return err
	}

	if err = s.db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}
