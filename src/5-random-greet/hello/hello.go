package main

import (
	"fmt"
	"log"

	"github.com/cirano-eusebi/go-starter/src/5-random-greet/greetings"
)

func main() {
	// Set the properties of the predefined Logger, including
	// the log entry prefix and the flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("5-random-greet: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("GLaDOS")
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
