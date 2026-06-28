package handlers

import (
	"fmt"
	"to-do-app-in-go/db"
	tps "to-do-app-in-go/types"

	"encoding/json"
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
		var dbError tps.ErrorResponse
		dbError.Error = fmt.Sprintf("No record of task with ID %s found", c.Params("task"))

		return c.Status(400).JSON(dbError)
	}
}

func UpdateTask(c *fiber.Ctx) error {
	return c.SendString("Update task")
}

func NewTask(c *fiber.Ctx) error {

	//ToDo: Refactor into helpers

	taskReq := c.Body()

	// ToDo: Need to add request validation here

	taskReqStr := fmt.Sprintf("%s", taskReq)
	fmt.Printf("Recieved write request for task: \n %s", taskReqStr)

	var newTask tps.TaskRequest

	err := json.Unmarshal([]byte(taskReq), &newTask)

	if err != nil {
		var errorResponse tps.ErrorResponse
		fmt.Printf("Invalid format for new task. New tasks must contain taskBody and status only.")
		errorResponse.Error = "Unexpected error parsing new task, New tasks must contain taskBody and status only"
		return c.Status(422).JSON(errorResponse)
	}

	var newTaskResponse tps.WriteTaskResponse
	fmt.Printf("Create new task: %+v\n", newTask)
	err = db.DbCreateTask(newTask)
	if err != nil {
		newTaskResponse.Error = fmt.Sprintf("Failed to write to database with error: %s\n", err)
		return c.Status(503).JSON(newTaskResponse)
	}
	newTaskResponse.TaskId = "Successfully wrote new task" // Todo: change to id of task
	return c.Status(200).JSON(newTaskResponse)

}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
