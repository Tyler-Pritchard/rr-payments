package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"stripe-payment-service/handlers/payment_handler"
	"stripe-payment-service/handlers/refund_handler"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	// Load environment variables from .env only if running locally
	if os.Getenv("DOCKER_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found, falling back to system environment variables.")
		}
	}

	// Set the Stripe API key from the environment variable
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		log.Fatalf("Stripe API key not set in environment variables.")
	}
	stripe.Key = stripeKey

	log.Println("Initializing Stripe Payment Service with Test Mode...")

	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/charge", payment_handler.HandleCharge)
	mux.HandleFunc("/refund", refund_handler.HandleRefund)

	// Register the health check route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"UP"}`)
	})

	// Start the HTTP server
	log.Println("Server starting on port 8082...")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
