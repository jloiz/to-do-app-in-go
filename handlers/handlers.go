package handlers

import (
	"github.com/gofiber/fiber/v2"
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

//ToDo: Should I move todo app items to own package?
// After made all endpoints, jsonify the default and send

func GetAllTasks(c *fiber.Ctx) error {
	return c.SendString("All tasks")
}

func GetTask(c *fiber.Ctx) error {
	msg := "A single task " + c.Params("task");
	return c.SendString(msg);
}

func UpdateTask(c *fiber.Ctx) error {
	return c.SendString("Update task")
}

func NewTask(c *fiber.Ctx) error {
	return c.SendString("New Task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
