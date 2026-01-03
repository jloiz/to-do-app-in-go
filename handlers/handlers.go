package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/db"
	tps "to-do-app-in-go/types"
)

func Greeting() string {
	var greetingStr string
	greetingStr = "Welcome, getting your tasks."
	return greetingStr
}

func Initialise(c *fiber.Ctx) error {
	greeting := Greeting()
	return c.SendString(greeting)
}

func GetAllTasks(c *fiber.Ctx) error {

	// This one is broken
	queryStr := "SELECT * FROM tasks"
	result := db.DbRead(queryStr)
	fmt.Sprint(result)
	return c.SendString("All tasks")
}

func GetTask(c *fiber.Ctx) error {
	dbRes := db.DbReadRow(c.Params("task"))
	if dbRes.TaskId != "" {
		return c.JSON(dbRes)
	} else {
		// Define here so only use mem if an error 
		var  dbError tps.ErrorRes
		dbError.Error = fmt.Sprintf("No record of task with ID %s found", c.Params("task"))

		return c.Status(400).JSON(dbError)
	}
}

func UpdateTask(c *fiber.Ctx) error {
	return c.SendString("Update task")
}

func NewTask(c *fiber.Ctx) error {
	// Use UUID
	return c.SendString("New Task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
