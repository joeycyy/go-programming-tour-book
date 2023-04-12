package main

import (
	"log"

	"github.com/go-programming-tour-book/tour/cmd"
)

func main() {
	log.Print("main")
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
