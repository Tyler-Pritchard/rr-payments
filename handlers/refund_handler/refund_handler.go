package refund_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/refund"
)

// RefundRequest represents the expected fields for the refund
type RefundRequest struct {
	ChargeID string `json:"charge_id"` // The ID of the charge to refund
}

// HandleRefund processes a refund request via Stripe
func HandleRefund(w http.ResponseWriter, r *http.Request) {
	var req RefundRequest

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request data
	if req.ChargeID == "" {
		log.Printf("Validation error: missing charge_id")
		http.Error(w, "Missing charge_id", http.StatusBadRequest)
		return
	}

	log.Printf("Processing refund for charge: %s", req.ChargeID)

	// Create the Stripe refund
	params := &stripe.RefundParams{
		Charge: stripe.String(req.ChargeID),
	}

	ref, err := refund.New(params)
	if err != nil {
		log.Printf("Refund failed: %v\n", err)
		http.Error(w, "Refund failed", http.StatusInternalServerError)
		return
	}

	log.Printf("Refund succeeded: %s, Amount: %d, Status: %s\n",
		ref.ID, ref.Amount, ref.Status)

	// Respond with the refund details
	response := map[string]interface{}{
		"refund_id": ref.ID,
		"amount":    ref.Amount,
		"status":    ref.Status,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
	}
}
