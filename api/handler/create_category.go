package handler

import (
	voteservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/vote_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(ctx *fiber.Ctx) error {

	var input *types.CreateCategoryInput
	ctx.BodyParser(&input)

	id := ctx.Locals("id").(uint)

	voteservice := voteservice.NewVoteService()

	err := voteservice.CreateCategory(input, id)

	if err != nil {
		return err
	}

	return ctx.SendStatus(201)
}