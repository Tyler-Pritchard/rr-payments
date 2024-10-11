# rr-payments

## Overview

```rr-payments``` is a Go-based microservice designed for secure and scalable payment processing using the Stripe API. The service provides basic payment functionality, including creating charges and issuing refunds, with built-in input validation, structured error handling, and Stripe's best practices.

## Key Features

- <b>Stripe Payment Integration:</b> Seamlessly handles charge creation and refunds using Stripe's Go SDK.
- <b>Microservice Architecture:</b> Designed as a standalone microservice with modularity and extensibility in mind.
- <b>Input Validation & Error Handling:</b> Implements thorough validation for payment fields and detailed error responses.
- <b>Well-Documented Code:</b> Clear structure and comprehensive logging for easy debugging and maintenance.
- <b>Secure Environment Configuration:</b> Utilizes environment variables for secure management of sensitive data (e.g., Stripe API keys).

## File Structure

```
rr-payments/
│
├── cmd/
│   └── stripe-payment-service/
│       └── main.go                   # Entry point for the service
│
├── handlers/
│   └── payment_handler/
│       └── payment_handler.go         # Logic for handling /charge requests
│
│   └── refund_handler/
│       └── refund_handler.go          # Logic for handling /refund requests
│
├── config.go                          # Configuration settings for environment variables
├── payments.go                        # Payment-related structs and utilities
├── go.mod                             # Go module configuration
├── go.sum                             # Go dependencies checksum
```

## Getting Started

### Prerequisites

- <b>Go 1.16 or later</b> installed
- <b>Stripe Account</b> with Test Mode enabled
- <b>Stripe API Key</b> stored as an environment variable in ```.env```:

```
STRIPE_SECRET_KEY=sk_test_xxxxxxxxxxxxxxxx
```

### Installation

1. Clone the repository:
```
git clone https://github.com/yourusername/rr-payments.git
cd rr-payments
```
2. Install dependencies:
```
go mod tidy
```
3. Run the service locally:
```
go run cmd/stripe-payment-service/main.go
```
4. The service will be available at ```http://localhost:8080```.

### Endpoints

```/charge``` - <b>Create a Payment Charge</b>
</b>Method:</b> ```POST```
<b>Request Body:</b>
```
{
  "amount": 2000,
  "currency": "usd",
  "source": "tok_visa"
}
```
- <b>Response:</b>
```
{
  "charge_id": "ch_1IuR7X2eZvKYlo2C8z0Bv2Cj",
  "amount": 2000,
  "currency": "usd",
  "status": "succeeded",
  "receipt_url": "https://pay.stripe.com/receipts/payment/..."
}
```
```/refund``` <b>- Issue a Refund for a Charge</b>
- <b>Method:</b> ```POST```
- <b>Request Body:</b>
```
{
  "charge_id": "ch_1IuR7X2eZvKYlo2C8z0Bv2Cj"
}
```
- <b>Response:</b>
```
{
  "refund_id": "re_1IuS8H2eZvKYlo2C8z0Bv2Cf",
  "amount": 2000,
  "status": "succeeded"
}
```
## Design Considerations

### Security

- All API keys and sensitive information are stored in environment variables and never hard-coded.
- The service is built using the latest version of the Stripe SDK to ensure compatibility and security.
- Implements input validation to prevent common attacks (e.g., SQL injection).

### Modularity

- Organized as a microservice with isolated concerns for payments and refunds.
- Easily extensible to support additional payment functionalities, such as subscriptions and webhooks.

### Performance
- Lightweight service with low memory footprint, designed to handle high-throughput payment requests efficiently.

## Future Enhancements

- Implementing subscription-based payments.
- Adding support for webhooks to track charge updates.
- Introducing a front-end client for managing payments and viewing transaction history.

## Technical Stack

- <b>Language:</b> Go
- <b>Framework:</b> Standard Library with ```net/http```
- <b>Stripe SDK:</b> ```github.com/stripe/stripe-go/v72```
- <b>Environment Management:</b> ```github.com/joho/godotenv```

# License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or collaboration, please reach out via LinkedIn or check out more of my work on [GitHub](https://www.github.com/tyler-pritchard/rr-payments).

