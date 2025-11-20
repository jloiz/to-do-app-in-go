package routes

import (
	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Health
	app.Get("/", handlers.Initialise)
}
