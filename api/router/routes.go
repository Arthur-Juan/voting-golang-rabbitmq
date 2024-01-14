package router

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/api/handler"
	"github.com/gofiber/fiber/v2"
)

func defineRoutes(router *fiber.App) {

	v1 := router.Group("api/v1")
	{
		v1.Get("/hello", func(c *fiber.Ctx) error {
			return c.JSON("Hello")
		})

		v1.Post("/auth/register", handler.CreateAccountHandler)
		v1.Post("/auth/login", handler.LoginHandler)
	}
}
