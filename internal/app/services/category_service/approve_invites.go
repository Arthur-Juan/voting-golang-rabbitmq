package categoryservice

import (
	"errors"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"gorm.io/gorm"
)

func (s *CategoryService) ApproveInvite(user_id, category_id, target_id int, decision types.Status) error {

	var candidate_category *types.CandidateCategory

	var admin_category *types.CategoryAdmin

	if err := s.db.Where(&types.CategoryAdmin{UserId: uint(user_id), CategoryId: uint(category_id)}).First(&admin_category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("you cannot do this action")
		}
		return err
	}

	if err := s.db.Where(&types.CandidateCategory{UserId: uint(target_id), CategoryId: uint(category_id), Status: types.Pending}).First(&candidate_category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not found")
		}
		return err
	}

	if decision != types.Rejected && decision != types.Approved {
		return errors.New("invalid decision")
	}

	candidate_category.Status = decision

	s.db.Save(candidate_category)

	return nil
}
