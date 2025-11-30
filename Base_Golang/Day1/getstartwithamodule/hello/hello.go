package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	// Get a greeting message with a name and handle any errors
	message, err := greeting.Hello("Go Learner")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(message)
}