package main

import (
	"log"
	"os"

	"github.com/ostcar/battlesnake/snake"
)

func main() {
	if err := snake.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
