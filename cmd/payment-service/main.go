package main

import (
	"log"
	"os"

	"net/http"
	"stripe-payment-service/handlers"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set the Stripe API key from the environment variable
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		log.Fatalf("Stripe API key not set in environment variables.")
	}
	stripe.Key = stripeKey

	log.Println("Initializing Stripe Payment Service with Test Mode...")

	// Register the /charge route
	http.HandleFunc("/charge", handlers.HandleCharge)

	// Start the HTTP server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
