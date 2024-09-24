package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

// ChargeRequest represents the expected fields for the payment
type ChargeRequest struct {
	Amount   int64  `json:"amount"`   // Amount in cents
	Currency string `json:"currency"` // Currency (e.g., "usd")
	Source   string `json:"source"`   // Payment source (token or card)
}

// ValidateChargeRequest checks if the input is valid
func ValidateChargeRequest(req ChargeRequest) error {
	// Validate the amount (must be positive)
	if req.Amount <= 0 {
		return errors.New("invalid amount: must be greater than 0")
	}

	// Validate the currency (here we just check if it's not empty, but you could also check against a list of valid currency codes)
	if req.Currency == "" {
		return errors.New("invalid currency: must not be empty")
	}

	// Validate the source (must not be empty)
	if req.Source == "" {
		return errors.New("invalid source: must not be empty")
	}

	return nil
}

// HandleCharge processes a payment charge via Stripe
func HandleCharge(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req ChargeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request data
	if err := ValidateChargeRequest(req); err != nil {
		log.Printf("Validation error: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Processing payment: %d %s", req.Amount, req.Currency)

	// Create the Stripe charge
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

	// Respond with the charge ID
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"charge_id": ch.ID})
}
