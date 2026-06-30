package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	tps "to-do-app-in-go/types"
	hlp "to-do-app-in-go/helpers"
)

// Global database connection ptr
var dbConn *sql.DB

func ConnectToDb() bool {
	//ToDo: Close db conn, recover block needed
	variables := [5]string{"HOST", "PORT", "USER", "PASSWORD", "DBNAME"}
	envVarMap := make(map[string]string)

	for _, envVar := range variables {
		buff := hlp.GetEnvVariable(envVar)
		envVarMap[envVar] = buff
	}

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", envVarMap["HOST"], envVarMap["PORT"], envVarMap["USER"], envVarMap["PASSWORD"], envVarMap["DBNAME"])
	var err error
	dbConn, err = sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	// ToDo: Find where to close
	//defer db.Close()
	err = dbConn.Ping()
	var isConnected bool

	if err != nil {
		isConnected = false
		panic(err)
	} else {
		isConnected = true
	}

	log.Printf("Connected to database successfully\n")

	return isConnected
}

func dbWrite(command string) (string, error) {
	var msg string
	var writtenId string
	err := dbConn.QueryRow(command).Scan(&writtenId)

	if err != nil {

		msg = fmt.Sprintf("\nDatabase write failed: %v", err)
		log.Printf("%s", msg)
		return writtenId, err
	}

	log.Printf("New task write process successful with id: %s", writtenId)
	return writtenId, nil
}

func DbCreateTask(newTask tps.TaskRequest) (string, error) {
	//ToDo: Move uuid generation to auto key generation in table config for postgress DB
	writeCommand := fmt.Sprintf("INSERT into tasks (task_id, task_body, status) values (uuid_generate_v4(), '%s','%s') RETURNING task_id;", newTask.TaskBody, newTask.Status)
	log.Printf("%s \n", writeCommand)
	taskId, err := dbWrite(writeCommand)
	log.Printf("%+v", err)
	return taskId, err
}

func DbReadRow(taskId string) (tps.Task, error) {
	var task tps.Task
	var noRes tps.Task

	queryTemplate := "SELECT task_id, task_body, status FROM tasks WHERE task_id=$1"
	row := dbConn.QueryRow(queryTemplate, taskId)
	err := row.Scan(&task.TaskId, &task.TaskBody, &task.Status)

	switch err {
	case sql.ErrNoRows:
		log.Printf("No task with %s found \n", taskId)
		return noRes, err
	case nil:
		// Success case
		return task, err
	default:
		return noRes, err
	}
}

func DbRead(command string) *sql.Rows {
	rows, err := dbConn.Query(command)
	if err != nil {
		panic(err)
		// may need to rethink for not found for successful write
		// Or probably just set postgres col as primary key as more realistic tbh
	}
	defer fmt.Sprint(rows)
	return rows
}


