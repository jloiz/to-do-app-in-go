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
	queryStr := "SELECT * FROM tasks"
	result := db.DbRead(queryStr)
	fmt.Sprint(result)
	return c.SendString("All tasks")
}

func GetTask(c *fiber.Ctx) error {
	// ToDo: Do some http side checking of validity of id provided
	dbRes := db.DbReadRow(c.Params("task"))
	if dbRes.TaskId != 0 {
		return c.JSON(dbRes)
	} else {
		// Define here so only use mem if an error 
		var  dbError tps.ErrorRes
		dbError.Error = fmt.Sprintf("No record of task with ID %s found", c.Params("task"))

		return c.JSON(dbError)
	}

}

func UpdateTask(c *fiber.Ctx) error {
	return c.SendString("Update task")
}

func NewTask(c *fiber.Ctx) error {
	// This one need a select first to check the id does not
	// already exisit
	return c.SendString("New Task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
