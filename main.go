package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/db"
	"to-do-app-in-go/routes"
)

func main() {
	connectionStatus := db.ConnectToDb()
	fmt.Printf("Connection: %t \n", connectionStatus)

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")

}
