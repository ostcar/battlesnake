package snake

import (
	"fmt"
	"log"
	"net/http"
)

// Run starts the webserver.
func Run(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/start", startHandler)
	mux.HandleFunc("/move", moveHandler)
	mux.HandleFunc("/end", endHandler)

	log.Printf("Listen on: %s", addr)
	if err := http.ListenAndServe(addr, mux); err != http.ErrServerClosed {
		return fmt.Errorf("running webserver: %w", err)
	}
	return nil
}
