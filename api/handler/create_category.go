package handler

import (
	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(ctx *fiber.Ctx) error {

	var input *types.CreateCategoryInput
	ctx.BodyParser(&input)

	id := uint(ctx.Locals("userId").(uint64))

	categoryservice := categoryservice.NewCategoryService()

	err := categoryservice.CreateCategory(input, id)

	if err != nil {
		return err
	}

	return ctx.SendStatus(201)
}
