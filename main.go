package main

import (
	"fmt"

	"github.com/jloiz/to-do-app-in-go/greeting"
)

func main(){
	fmt.Println("Welcome,",  greeting.GreetUser())
}