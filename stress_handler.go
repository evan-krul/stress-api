package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type jsonBody struct {
	Args    []string `json:"args"`
	Timeout string   `json:"timeout"`
}

func StressHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	decoder := json.NewDecoder(r.Body)
	var body jsonBody
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	go func() {
		var command []string
		command = append(command, body.Args...)
		command = append(command, "--timeout", body.Timeout)

		cmd := exec.Command("stress-ng", command...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout

		// run command
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	_, err = io.WriteString(w, "Stress")
	if err != nil {
		return
	}
}
