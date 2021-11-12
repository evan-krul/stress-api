package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/stress", StressHandler).Methods("POST")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
