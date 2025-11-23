package main

import (

	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/routes"
)

func main() {

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")

}
