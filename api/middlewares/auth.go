package middlewares

import (
	"fmt"
	"net/http"

	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	fmt.Println(tokenString)
	if tokenString == "" {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"msg": "Forbidden",
		})
	}
	// Remove the "Bearer " prefix from the token
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := token.CheckToken(tokenString)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
	}
	ctx.Locals("userId", token.ID)

	return ctx.Next()
}
