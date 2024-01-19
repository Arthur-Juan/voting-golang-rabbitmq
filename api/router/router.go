package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Initialize() {
	router := fiber.New()
	// Initialize default config
	router.Use(recover.New())
	defineRoutes(router)

	err := router.Listen(":8080")
	if err != nil {
		return
	}
}
