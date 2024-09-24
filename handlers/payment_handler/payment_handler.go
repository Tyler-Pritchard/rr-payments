package payment_handler

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
	if req.Amount <= 0 {
		return errors.New("invalid amount: must be greater than 0")
	}
	if req.Currency == "" {
		return errors.New("invalid currency: must not be empty")
	}
	if req.Source == "" {
		return errors.New("invalid source: must not be empty")
	}
	return nil
}

// HandleCharge processes a payment charge via Stripe
func HandleCharge(w http.ResponseWriter, r *http.Request) {
	var req ChargeRequest

	// Parse the request body
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
		if stripeErr, ok := err.(*stripe.Error); ok {
			// Handle different types of Stripe errors
			switch stripeErr.Code {
			case stripe.ErrorCodeCardDeclined:
				log.Printf("Charge failed: Card Declined - %v\n", stripeErr)
				http.Error(w, "Card was declined", http.StatusPaymentRequired)
			case stripe.ErrorCodeExpiredCard:
				log.Printf("Charge failed: Expired Card - %v\n", stripeErr)
				http.Error(w, "Card has expired", http.StatusPaymentRequired)
			case stripe.ErrorCodeIncorrectCVC:
				log.Printf("Charge failed: Incorrect CVC - %v\n", stripeErr)
				http.Error(w, "Incorrect CVC", http.StatusPaymentRequired)
			default:
				// Generic error handling for invalid requests
				if _, isInvalidRequest := err.(*stripe.InvalidRequestError); isInvalidRequest {
					log.Printf("Charge failed: Invalid Request - %v\n", stripeErr)
					http.Error(w, "Invalid request parameters", http.StatusBadRequest)
				} else {
					log.Printf("Charge failed: %v\n", stripeErr)
					http.Error(w, "Payment processing error", http.StatusInternalServerError)
				}
			}
		} else {
			// Handle generic errors
			log.Printf("Charge failed: %v\n", err)
			http.Error(w, "Payment processing error", http.StatusInternalServerError)
		}
		return
	}

	log.Printf("Charge succeeded: %s, Amount: %d %s, Status: %s, Receipt: %s\n",
		ch.ID, ch.Amount, ch.Currency, ch.Status, ch.ReceiptURL)

	// Respond with the relevant charge details
	response := map[string]interface{}{
		"charge_id":   ch.ID,
		"amount":      ch.Amount,
		"currency":    ch.Currency,
		"status":      ch.Status,
		"receipt_url": ch.ReceiptURL,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
	}
}
