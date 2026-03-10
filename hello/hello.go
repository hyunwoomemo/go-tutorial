package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"hyunwoo", "eunbi", "jian"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	for _, message := range messages {
		fmt.Println(message)
	}

	message, err := greetings.Hello("hw")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

