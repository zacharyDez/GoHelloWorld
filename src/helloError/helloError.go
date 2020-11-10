package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined loggers
	// includes the log entry prefix 
	// and a flag to disable printing, 
	// time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	message, err := greetings.HelloError("Glady's")

	// if an error is returned, print it to console 
	// and quit
	if err != nil {
		log.Fatal(err)
	}
	
	// print message
	fmt.Println(message)
}

