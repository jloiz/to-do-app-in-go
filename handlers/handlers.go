package handlers

import(
	"to-do-app-in-go/greeting"
	"github.com/gofiber/fiber/v2"

)

func MakeGreeting() string {
	user := greeting.GreetUser()
	var greetingStr string
	greetingStr = "Welcome " + user + ", getting your tasks."
	return greetingStr
}

// ToDo: Move route handlers to their own gomod
func Initialise(c *fiber.Ctx) error {
	greeting := MakeGreeting()
	return c.SendString(greeting)
}