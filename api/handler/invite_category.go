package handler

import (
	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func InviteToCategoryHandler(ctx *fiber.Ctx) error {

	var input *types.InviteToCategoryInput
	ctx.BodyParser(&input)

	id := uint(ctx.Locals("userId").(uint64))

	service := categoryservice.NewCategoryService()
	err := service.InviteToCategory(input, id)

	if err != nil {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(201)
}
