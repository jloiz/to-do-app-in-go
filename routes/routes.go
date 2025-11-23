package routes

import (
	"to-do-app-in-go/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health
	app.Get("/", handlers.Initialise)
	app.Get("/demo", handlers.Initialise2)

}
