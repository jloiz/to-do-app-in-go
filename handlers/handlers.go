package handlers

import (
	"fmt"
	"strings"

	"to-do-app-in-go/db"
	"to-do-app-in-go/utils"

	tps "to-do-app-in-go/types"

	"github.com/gofiber/fiber/v2"
)

// Globally available request status constant

// Common vars
var (
	newTask         tps.TaskRequest
	taskId          string
	newTaskResponse tps.WriteTaskResponse
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

////// GETs  //////

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

func GetAllTasks(c *fiber.Ctx) error {

	// This one is broken
	queryStr := "SELECT * FROM tasks"
	result := db.DbRead(queryStr)
	fmt.Sprint(result)
	return c.SendString("All tasks")
}

////// POSTs  //////

func NewTask(c *fiber.Ctx) error {
	taskReq := c.Body()

	newTask, err := utils.ParseRequest(taskReq)
	if err != nil {
		fmt.Printf("Unexpected error parsing new task, New tasks must contain taskBody and status only: %v", err)
		return c.Status(422).JSON(tps.ErrorResponse{
			Error: fmt.Sprintf("Unexpected error parsing new task, New tasks must contain taskBody and status only: %v", err),
		})
	}

	fmt.Printf("Create new task: %+v\n", newTask)

	// Up case status field for DB consistency
	newTask.Status = strings.ToUpper(newTask.Status)

	err = utils.ValidateRequest(newTask)
	if err != nil {
		fmt.Printf("Error validating request: %v", err)
		return c.Status(400).JSON(tps.ErrorResponse{
			Error: fmt.Sprintf("Error validating request: %v", err),
		})
	}

	// Write to db
	taskId, err = db.DbCreateTask(newTask)

	if err != nil {
		fmt.Printf("Error performing write to database: %v", err)
		return c.Status(503).JSON(tps.ErrorResponse{
			Error: fmt.Sprintf("Error performing write to database: %v", err),
		})
	}

	newTaskResponse.TaskId = taskId
	return c.Status(200).JSON(newTaskResponse)

}

func UpdateTask(c *fiber.Ctx) error {

	newTaskReq := c.Body()

	newTaskReqStr := fmt.Sprintf("%s", newTaskReq)
	fmt.Printf("Recieved write request for task: \n %s\n", newTaskReqStr)
	return c.SendString("Update task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
