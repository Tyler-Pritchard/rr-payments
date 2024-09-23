package config

import (
	"log"
	"os"
)

func GetStripeKey() string {
	log.Println("Fetching Stripe Secret Key from environment variables.")

	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Println("Warning: STRIPE_SECRET_KEY is not set.")
	} else {
		log.Println("Stripe Secret Key fetched successfully.")
	}

	return key
}
