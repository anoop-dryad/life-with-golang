package main

import (
	"fmt"
	"log"

	"godev.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Anoop")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
