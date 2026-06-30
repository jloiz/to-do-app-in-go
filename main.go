package main

import (
	"fmt"
	"os"
	"log"

	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/db"
	"to-do-app-in-go/routes"
)

func main() {
	log.SetOutput(os.Stdout)

	connectionStatus := db.ConnectToDb()
	fmt.Printf("Connection: %t \n", connectionStatus)

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")

}
