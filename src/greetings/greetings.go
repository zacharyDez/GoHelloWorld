package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Hello returns a greeting for a named person
func HelloError(name string) (string, error) {
	// no name provided, throw an error
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name provided, embed the name 
	// in the greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
