package router

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/api/handler"
	"github.com/arthur-juan/voting-golang-rabbitmq/api/middlewares"
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	jwtware "github.com/gofiber/contrib/jwt"
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

	auth := v1.Group("", middlewares.CheckAuth)
	{
		auth.Use(jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{Key: []byte(config.GetKey())}}))

		auth.Post("/category", handler.CreateCategory)
		auth.Get("/category", handler.ListCategory)
	}
}
