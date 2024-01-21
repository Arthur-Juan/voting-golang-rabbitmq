package handler

import (
	"fmt"

	voteservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/vote_service"
	"github.com/gofiber/fiber/v2"
)

func ListCategory(ctx *fiber.Ctx) error {
	fmt.Println("HEEEERE")
	service := voteservice.NewVoteService()
	fmt.Println(service)
	output, err := service.ListCategory()
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(output)

}
