package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/start", startHandler)
	mux.HandleFunc("/move", moveHandler)
	mux.HandleFunc("/end", endHandler)

	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		return fmt.Errorf("running webserver: %w", err)
	}
	return nil
}
