package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FindInArray(chkVal interface{}, chkArr []interface{}) int {
	// Checks if a value is in an array and returns the position if present
	// Returns -1 if the value is not found in the array
	for i, v := range chkArr {
		if v == chkVal {
			return i
		}
	}
	return -1
}

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to fetch environment variable: %s", key)
	}
	return os.Getenv(key)
}
