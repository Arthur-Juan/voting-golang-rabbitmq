package handler

import (
	"fmt"

	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/gofiber/fiber/v2"
)

func ListCategory(ctx *fiber.Ctx) error {
	service := categoryservice.NewCategoryService()
	fmt.Println(service)
	output, err := service.ListCategory()
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(output)

}
