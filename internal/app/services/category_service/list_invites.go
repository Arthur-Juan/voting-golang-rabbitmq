package categoryservice

import (
	"errors"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/gorm"
)

func (s *CategoryService) ListInvites(user_id, category_id uint) ([]types.CandidateCategoryOutput, error) {

	var admin_category *types.CategoryAdmin

	if err := s.db.Where(&types.CategoryAdmin{UserId: user_id, CategoryId: category_id}).First(&admin_category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("you not have permission to see the invites of this category")
		}
	}

	var candidateCategory []*types.CandidateCategory

	if err := s.db.Preload("User").Preload("Category").
		Find(&candidateCategory, "category_id = ?", category_id).
		Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}

		return nil, err
	}

	var output []types.CandidateCategoryOutput

	for _, value := range candidateCategory {
		output = append(output, types.CandidateCategoryOutput{
			User: &value.User,
			Category: &types.CategoryOutput{
				ID:          value.CategoryId,
				Name:        value.Category.Name,
				Description: value.Category.Description,
				End:         value.Category.End,
				Winners:     int(value.Category.Winners),
				Status:      "status",
			},
		})
	}

	return output, nil
}
