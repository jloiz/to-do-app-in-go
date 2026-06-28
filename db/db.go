package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_"github.com/lib/pq"

	tps "to-do-app-in-go/types"
)

// Global database connection ptr
var dbConn *sql.DB

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to fetch environment variable: %s", key)
	}

	return os.Getenv(key)
}

func ConnectToDb() bool {
	//ToDo: Close db conn, recover block needed
	variables := [5]string{"HOST", "PORT", "USER", "PASSWORD", "DBNAME"}
	envVarMap := make(map[string]string)

	for _, envVar := range variables {
		buff := getEnvVariable(envVar)
		envVarMap[envVar] = buff
	}

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", envVarMap["HOST"], envVarMap["PORT"], envVarMap["USER"], envVarMap["PASSWORD"], envVarMap["DBNAME"])
	var err error;
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

	fmt.Printf("Connected to database successfully\n")

	return isConnected
}

func dbWrite(command string) error {
	 // Define the recovery block for bad database writes
	 var msg string;

	resp, err := dbConn.Exec(command)
	if err != nil {
		
		msg = fmt.Sprintf("\nDatabase write failed: %v", err)
		fmt.Printf("%s", msg)
		return err
	}
	fmt.Printf("%s", resp)
	// ToDo: Change to switch statement and get data from post
	msg = "New task write process successful"
	fmt.Printf("%s", msg)
	return nil
}

func DbCreateTask(newTask tps.TaskRequest) error {
	//ToDo: Move uuid generation to auto key generation in table config for postgress DB
	writeCommand := fmt.Sprintf("INSERT into tasks (task_id, task_body, status) values (uuid_generate_v4(), '%s','%s') RETURNING task_id;", newTask.TaskBody, newTask.Status)
	fmt.Printf("%s \n", writeCommand)
	err := dbWrite(writeCommand)
	fmt.Printf("%+v", err);
	return err
}

func DbReadRow(taskId string) tps.Task {
	var task tps.Task
	var noRes tps.Task

	queryTemplate := "SELECT task_id, task_body, status FROM tasks WHERE task_id=$1"
	row := dbConn.QueryRow(queryTemplate, taskId)
	err := row.Scan(&task.TaskId, &task.TaskBody, &task.Status)
    
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("No task with %s found \n", taskId)
		return noRes
	case nil: 
		// Success case
		return task
	default:
		panic(err)
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


	//  defer func() {
	// 	if r:= recover(); r != nil {
	// 		msg = fmt.Sprintf("\nDatabase write failed: %v", r)
	// 		fmt.Printf("%s", msg)
	// 	}
	//  }() 