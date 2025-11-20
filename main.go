package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jloiz/to-do-app-in-go/greeting"
)

func makeGreeting() string {
	user := greeting.GreetUser()
	var greetingStr string
	greetingStr = "Welcome " + user + ", getting your tasks."
	return greetingStr
}

func setupRoutes(app *fiber.App) {
	// Health 
	app.Get("/", initialise)
}

// ToDo: Move route handlers to their own gomod
func initialise(c *fiber.Ctx) error {
	greeting := makeGreeting()
	return c.SendString(greeting)
}

func main() {

	greeting := makeGreeting()
	fmt.Println(greeting)

	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")

}
