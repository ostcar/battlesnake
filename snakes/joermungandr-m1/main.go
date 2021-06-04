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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving info request")
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{
		"apiversion": "1",
		"author": "Oskar Hahn",
		"color" : "#123456",
		"head" : "default",
		"tail" : "default",
		"version" : "0.0.1-beta"
	}`))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving start request")
	w.Header().Set("Content-Type", "application/json")
	// TODO Start game
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving move request")
	w.Header().Set("Content-Type", "application/json")
	// TODO find game, move and return

	w.Write([]byte(`{
		"move": "up",
		"shout": "I am moving up!"
	  }`))
}

func endHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving end request")
	w.Header().Set("Content-Type", "application/json")
	// Remove game from store
}
