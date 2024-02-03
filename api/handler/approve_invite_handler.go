package handler

import (
	categoryservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/category_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func ApproveInviteHandler(ctx *fiber.Ctx) error {
	var input *struct {
		TargetId   int          `json:"target_id"`
		CategoryId int          `json:"category_id"`
		Choice     types.Status `json:"choice"`
	}

	id := int(ctx.Locals("userId").(uint64))
	ctx.BodyParser(&input)

	svc := categoryservice.NewCategoryService()

	err := svc.ApproveInvite(id, input.CategoryId, input.TargetId, input.Choice)

	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
