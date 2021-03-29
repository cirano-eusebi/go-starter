package main

import (
	"fmt"

	"github.com/cirano-eusebi/go-starter/src/3-create-module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("GlaDOS")
	fmt.Println(message)
}
