package main

import (
	"log"
	"os"

	"github.com/ostcar/battlesnake/snake"
)

func main() {
	// TODO: Use flags, also for debug and lookahead
	addr := ":8080"
	if len(os.Args) >= 2 {
		addr = os.Args[1]
	}

	if err := snake.Run(addr); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
