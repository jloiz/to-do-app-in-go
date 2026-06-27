package handlers

import (
	"fmt"
	"to-do-app-in-go/db"
	tps "to-do-app-in-go/types"

	"github.com/gofiber/fiber/v2"
	"encoding/json"
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
	taskReq := c.Body()
	taskReqStr := fmt.Sprintf("%s", taskReq)
	fmt.Printf("Recieved write request for task: \n %s", taskReqStr)
	// ToDo: Got here, JSONisfy and validate type conformity. Throw err is not matched
	var resp string;
	
	var newTask tps.TaskRequest;
	err := json.Unmarshal([]byte(taskReq), &newTask) 
	if err == nil {
		fmt.Printf("Create new task: %+v\n", newTask)
		resp = "Success"
		// db.DbCreateTask(newTask);
	}
	if err != nil {
		//panic("Error recieving JSON from task")
		fmt.Printf("Invalid format for new task. New tasks must contain taskBody and status only.")
		// Change to proper respnse types with err
		resp = "failure"		
	}
	return c.SendString(resp)
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}
