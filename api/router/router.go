package router

import "github.com/gofiber/fiber/v2"

func Initialize() {
	router := fiber.New()

	defineRoutes(router)

	err := router.Listen(":8080")
	if err != nil {
		return
	}
}
