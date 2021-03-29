package main

import (
	"fmt"
	"log"

	"github.com/cirano-eusebi/go-starter/src/6-multi-greeter/greetings"
)

func main() {
	log.SetPrefix("6-multi-greeter: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{
		"GLaDOS",
		"Companion Cube",
		"Turret",
	}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console
	//fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}
