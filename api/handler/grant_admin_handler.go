package handler

import (
	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func GrantAdminHandler(ctx *fiber.Ctx) error {
	var input *types.GrantAdminAccessInput
	ctx.BodyParser(&input)

	user_id := uint(ctx.Locals("userId").(uint64))

	svc := categoryservice.NewCategoryService()
	err := svc.GrantAdminAccess(user_id, input)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
