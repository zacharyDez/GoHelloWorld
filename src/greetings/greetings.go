package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


// Hellos returns a map that associates each 
// of the name people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)

	// Loop though the received slice of names, calling 
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := HelloError(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. 
// The returned message is selected at random
func randomFormat() string {
	// a clice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format 
	return formats[rand.Intn(len(formats))]

}
