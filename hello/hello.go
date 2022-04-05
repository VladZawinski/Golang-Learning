package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main()  {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Thiha")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}