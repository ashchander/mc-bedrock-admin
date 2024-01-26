package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// dummy Get path to start out with
	r.HandleFunc("/get-commands", getHandler).Methods("GET")
	r.HandleFunc("/run-commands", runHandler).Methods("POST")

	// Set the static directory, and serve files
	staticFileDirectory := http.Dir("./public/")
	staticFileHandler := http.FileServer(staticFileDirectory)

	// Send anything at the root path to the static file directory
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	// Serve on 8080
	http.ListenAndServe(":8080", r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type CmdRequest struct {
Command string `json:"command"`
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var cmd CmdRequest;
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &cmd)
	if cmd.Command == "showcoordinates" {
		fmt.Println("executing showing coordinates")
	} else if cmd.Command == "hidecoordinates" {
		fmt.Println("executing hiding coordinates")
	}
	fmt.Fprintf(w, "{\"text\": \"Saved!\"}")
}
