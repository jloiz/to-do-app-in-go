package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"to-do-app-in-go/db"
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

// After made all endpoints, jsonify the default and send

func GetAllTasks(c *fiber.Ctx) error {
	queryStr := "SELECT * FROM tasks"
	result := db.DbRead(queryStr)
	fmt.Sprint(result)
	return c.SendString("All tasks")

}

func GetTask(c *fiber.Ctx) error {
	// Do some http side checking of validity of id provided
	db.DbReadRow(c.Params("task"))
	msg := "A single task " + c.Params("task")
	return c.JSON(msg)
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
