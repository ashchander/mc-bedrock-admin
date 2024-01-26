package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"

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
	var err error
	var cmd *exec.Cmd

	// Trigger Minecraft server to list game rules
	echo := "echo \"gamerule\" > /run/minecraft.stdin"
	cmd = exec.Command("bash", "-c", echo)
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "An error ocurred while trying to send the gamerule command")
		return
	}

	time.Sleep(time.Second)
	// Retrieve the output of the gamerule command
	journal := "journalctl -u minecraft.service -n 1 --no-pager"
	cmd = exec.Command("bash", "-c", journal)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "An error ocurred while trying to retrieve the gamerule command output")
		return
	}
	fmt.Println(string(stdout))
	fmt.Fprintf(w, string(stdout))
}

type CmdRequest struct {
Command string `json:"command"`
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var cmd *exec.Cmd
	var cmdRequest CmdRequest
	var cmdText string
	
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &cmdRequest)
	if cmdRequest.Command == "showcoordinates" {
		fmt.Println("executing showing coordinates")
		cmdText = "echo \"gamerule showcoordinates true\" > /run/minecraft.stdin"
		cmd = exec.Command("bash", "-c", cmdText)
		err = cmd.Run()
	} else if cmdRequest.Command == "hidecoordinates" {
		fmt.Println("executing hiding coordinates")
		cmdText = "echo \"gamerule showcoordinates false\" > /run/minecraft.stdin"
		cmd = exec.Command("bash", "-c", cmdText)
		err = cmd.Run()
	}
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "An error ocurred while updating the settings")
		return

	}
	fmt.Fprintf(w, "{\"text\": \"Saved!\"}")
}
