package handler

import (
	authservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/auth_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

func CreateAccountHandler(ctx *fiber.Ctx) error {

	var input *types.CreateAccountInput
	ctx.BodyParser(&input)

	authservice := authservice.NewAuthService()
	token, err := authservice.CreateAccount(input)
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]string{
		"token": token,
	})
}
