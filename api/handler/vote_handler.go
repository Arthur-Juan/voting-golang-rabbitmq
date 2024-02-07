package handler

import (
	voteservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/vote_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func VoteHandler(ctx *fiber.Ctx) error {

	var input *types.VoteInput
	ctx.BodyParser(&input)

	svc := voteservice.NewVoteService()
	id := uint(ctx.Locals("userId").(uint64))

	err := svc.Vote(id, input)
	if err != nil {
		return err
	}

	return nil
}
