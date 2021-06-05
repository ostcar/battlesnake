package snake

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		handleError(w, err)
		return
	}

	d := ai(p)

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
