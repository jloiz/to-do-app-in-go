package routes

import (
	"to-do-app-in-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health
	app.Get("/", handlers.Initialise)

	app.Get("/api/v1/todo/get_all", handlers.GetAllTasks)
	app.Get("/api/v1/todo/get/:task", handlers.GetTask)
	app.Post("/api/v1/todo/update/:task", handlers.UpdateTask)
	app.Post("/api/v1/todo/new", handlers.NewTask)
	app.Delete("/api/v1/todo/delete/:task", handlers.DeleteTask)
}
