package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// dummy Get path to start out with
	r.HandleFunc("/hello", handler).Methods("GET")

	// Set the static directory, and serve files
	staticFileDirectory := http.Dir("./public/")
	staticFileHandler := http.FileServer(staticFileDirectory)

	// Send anything at the root path to the static file directory
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	// Serve on 8080
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
