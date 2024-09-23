package handlers

import (
	"log"
	"net/http"
)

func HandlePayment(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request at /payment endpoint.")

	// You could log the HTTP method for debugging purposes
	log.Printf("Request method: %s\n", r.Method)

	// Log some placeholder processing message
	log.Println("Processing payment...")

	// Send a response to the client
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Payment Handler Reached!"))
	if err != nil {
		log.Printf("Error sending response: %v\n", err)
	}

	log.Println("Payment processed successfully.")
}
