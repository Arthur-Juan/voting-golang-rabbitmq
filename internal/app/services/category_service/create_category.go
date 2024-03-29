package categoryservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/worker"
	"github.com/gofiber/fiber/v2/log"
)

func (s *CategoryService) CreateCategory(input *types.CreateCategoryInput, user_id uint) error {
	var user *types.User

	if err := s.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return err
	}

	category, err := types.NewCategory(input.Name, input.Description, input.End, user, input.Winners)
	if err != nil {
		return err
	}

	if err := s.db.Create(&category).Error; err != nil {
		return err
	}

	if err := s.db.Preload("Admins").First(&category, category.ID).Error; err != nil {
		return err
	}

	admin_category := types.CategoryAdmin{
		UserId:     user.ID,
		CategoryId: category.ID,
	}

	if err := s.db.Create(&admin_category).Error; err != nil {
		return err
	}

	log.Debug(category)
	//start worker
	go worker.Run(category.End)

	return nil
}
