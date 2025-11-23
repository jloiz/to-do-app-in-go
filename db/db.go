package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Unable to fetch environment variable: %s", key)
	}

	return os.Getenv(key)
}

func ConnectToDb() string {

	statusStr := getEnvVariable("DBNAME")


	return statusStr
}