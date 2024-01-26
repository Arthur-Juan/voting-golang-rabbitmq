package categoryservice

import (
	"errors"
	"time"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/gorm"
)

func (s *CategoryService) InviteToCategory(input *types.InviteToCategoryInput, user_id uint) error {

	var category *types.Category
	if err := s.db.Where("id = ?", input.CategoryId).Find(&category).Error; err != nil {
		return errors.New("category not found")
	}

	if category.End.Before(time.Now()) {
		return errors.New("category already ended")
	}

	var candidateCategory types.CandidateCategory
	result := s.db.Where(&types.CandidateCategory{UserId: user_id, CategoryId: input.CategoryId}).First(&candidateCategory)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		} else {
			return result.Error
		}
	}

	if candidateCategory.UserId != 0 {
		return errors.New("user already invited for this category")
	}

	newCandidateCategory := &types.CandidateCategory{
		UserId:     user_id,
		CategoryId: input.CategoryId,
		Status:     types.Pending,
	}

	result = s.db.Create(newCandidateCategory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
