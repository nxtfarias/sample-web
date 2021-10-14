package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// First:
	//fmt.Println(quote.Go())

	// Second: Using the module greeting
	// Get a greeting message and print it
	//message := greetings.Hello("Gladys")
	//fmt.Println(message)

	// Third: Managing errors
	// Set the log properties including log entry prefix and a flag to disable printing
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Requesting a greeting message:
	//message, err := greetings.Hello("Gladys")

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	// if an error is returned, as we can expect, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error, print the return message
	//fmt.Println(message)
	fmt.Println(messages)
}
