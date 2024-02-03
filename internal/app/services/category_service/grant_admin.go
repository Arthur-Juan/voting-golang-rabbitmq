package categoryservice

import (
	"errors"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/gorm"
)

func (s *CategoryService) GrantAdminAccess(user_id uint, input *types.GrantAdminAccessInput) error {

	//check permission
	var admin_category *types.CategoryAdmin
	if err := s.db.Where(&types.CategoryAdmin{UserId: uint(user_id), CategoryId: uint(input.CategoryId)}).First(&admin_category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("you cannot do this action")
		}
		return err
	}

	//check if user is not candidate on the selected category
	var candidates *types.CandidateCategory
	if err := s.db.Where(&types.CandidateCategory{UserId: uint(input.TargetId), CategoryId: uint(input.CategoryId)}).First(&candidates).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user is a candidate on this category")
		}

	}

	new_admin := types.CategoryAdmin{
		UserId:     uint(input.TargetId),
		CategoryId: uint(input.CategoryId),
	}

	result := s.db.Create(&new_admin)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
