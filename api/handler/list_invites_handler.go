package handler

import (
	"fmt"
	"strconv"

	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/gofiber/fiber/v2"
)

func ListInvitesHandler(ctx *fiber.Ctx) error {

	service := categoryservice.NewCategoryService()

	user_id := uint(ctx.Locals("userId").(uint64))
	category_id_str := ctx.Params("category_id")
	fmt.Println("aqui 0")
	category_id_uint, err := strconv.ParseUint(category_id_str, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category_id",
		})
	}
	fmt.Println("aqui 1")

	output, err := service.ListInvites(user_id, uint(category_id_uint))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list invites",
		})
	}
	fmt.Println("aqui 2")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": output,
	})
}
