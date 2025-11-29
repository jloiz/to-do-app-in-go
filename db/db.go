package db

import (
	"log"
	"os"
	"fmt"
	"database/sql"
	_"github.com/lib/pq"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Unable to fetch environment variable: %s", key)
	}

	return os.Getenv(key)
}

func ConnectToDb() bool {
	variables := [5]string{"HOST", "PORT", "USER", "PASSWORD", "DBNAME"}
	envVarMap := make(map[string]string)
	
	for _,envVar := range variables {
		buff := getEnvVariable(envVar)
		envVarMap[envVar] = buff
	}

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", envVarMap["HOST"], envVarMap["PORT"], envVarMap["USER"], envVarMap["PASSWORD"], envVarMap["DBNAME"])  

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	var isConnected bool;

	if err != nil {
		panic(err)
		isConnected = false
	} else {
		isConnected = true
	}

	fmt.Printf("Connected to database successfully\n")

	return isConnected 
}
