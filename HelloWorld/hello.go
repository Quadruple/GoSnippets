package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message, err := greetings.Hello("Atakan Ataman")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Atakan", "Ataman", "Atik"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(messages["Atakan"])
}
