package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

type ChargeRequest struct {
	Amount   int64  `json:"amount"`   // Amount in cents (e.g., 2000 for $20.00)
	Currency string `json:"currency"` // Currency (e.g., "usd")
	Source   string `json:"source"`   // Payment source (token or card)
}

// HandleCharge processes a payment charge via Stripe
func HandleCharge(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req ChargeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Processing payment: %d %s", req.Amount, req.Currency)

	// Create a Stripe charge
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
		Source:   &stripe.SourceParams{Token: stripe.String(req.Source)},
	}

	ch, err := charge.New(params)
	if err != nil {
		log.Printf("Charge failed: %v\n", err)
		http.Error(w, "Charge failed", http.StatusInternalServerError)
		return
	}

	log.Printf("Charge succeeded: %s\n", ch.ID)

	// Respond with the charge ID as a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"charge_id": ch.ID})
}
