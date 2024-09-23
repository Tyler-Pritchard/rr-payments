package payments

import (
	"log"
)

func CreatePayment(amount int64, currency string) {
	log.Printf("Creating payment of %d %s\n", amount, currency)

	// Simulate a delay in processing (you could add real logic later)
	log.Println("Connecting to payment gateway...")

	// Simulate success
	log.Println("Payment creation successful.")
}
