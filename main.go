package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"to-do-app-in-go/handlers"
	"to-do-app-in-go/routes"
)

func main() {

	greeting := handlers.MakeGreeting()
	fmt.Println(greeting)

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")

}
