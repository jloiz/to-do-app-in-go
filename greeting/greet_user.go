package greeting

import (
	"fmt"
)

func GreetUser() string {
	// ToDo: Get User name and format the below
	var userName string
	fmt.Println("Enter user name:")
	fmt.Scan(&userName)

	return userName
}