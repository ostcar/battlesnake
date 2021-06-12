package snake

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Printf("Start game")
	w.Header().Set("Content-Type", "application/json")
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, fmt.Errorf("reading request body: %w", err))
	}

	var p payload
	if err := json.Unmarshal(buf, &p); err != nil {
		handleError(w, fmt.Errorf("decoding request body\nBody: %s\nError: %w", buf, err))
		return
	}

	d := ai(stateFromPayload(p))

	fmt.Fprintf(
		w,
		`{
			"move": "%s",
			"shout": "I will get you!"
		}`,
		d,
	)
}

func endHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("Error: %v", err)
	http.Error(w, "Ups, something went wrong :(", 500)
}
