package main

import (
	"fmt"

	"github.com/jloiz/to-do-app-in-go/greeting"
	"github.com/gofiber/fiber/v2"
)

func main(){

	user := greeting.GreetUser()

	var greetingStr string;
	
	greetingStr =  "Welcome " +  user + ", getting your tasks."

	fmt.Println(greetingStr)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(greetingStr)
	})

	app.Listen(":3000")

}