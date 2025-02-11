package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents a simple JSON response
type Response struct {
	Message string `json:"message"`
}

// handler function for Service A
func handlerA(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Service A!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handler function for Service B
func handlerB(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Service B!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	go func() {
		http.HandleFunc("/service-a", handlerA)
		log.Println("Starting Service A on :8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Service A failed to start: ", err)
		}
	}()

	http.HandleFunc("/service-b", handlerB)
	log.Println("Starting Service B on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Service B failed to start: ", err)
	}
}
