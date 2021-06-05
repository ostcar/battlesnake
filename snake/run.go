package snake

import (
	"fmt"
	"net/http"
)

// Run starts the webserver.
func Run() error {
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
