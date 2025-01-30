# rr-payments

## Overview

```rr-payments``` is a Go-based microservice designed for secure and scalable payment processing using the Stripe API. The service provides basic payment functionality, including creating charges and issuing refunds, with built-in input validation, structured error handling, and Stripe's best practices.

## Key Features

- **Stripe Payment Integration:** Seamlessly handles charge creation and refunds using Stripe's Go SDK.
- **Microservice Architecture:** Designed as a standalone microservice with modularity and extensibility in mind.
- **Input Validation & Error Handling:** Implements thorough validation for payment fields and detailed error responses.
- **Well-Documented Code:** Clear structure and comprehensive logging for easy debugging and maintenance.
- **Secure Environment Configuration:** Utilizes environment variables for secure management of sensitive data (e.g., Stripe API keys).

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

- **Go 1.16 or later** installed
- **Stripe Account** with Test Mode enabled
- **Stripe API Key** stored as an environment variable in `.env`:

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
4. The service will be available at `http://localhost:8080`.

### Endpoints

#### `/charge` - **Create a Payment Charge**
- **Method:** `POST`
- **Request Body:**
```
{
  "amount": 2000,
  "currency": "usd",
  "source": "tok_visa"
}
```
- **Response:**
```
{
  "charge_id": "ch_1IuR7X2eZvKYlo2C8z0Bv2Cj",
  "amount": 2000,
  "currency": "usd",
  "status": "succeeded",
  "receipt_url": "https://pay.stripe.com/receipts/payment/..."
}
```

#### `/refund` - **Issue a Refund for a Charge**
- **Method:** `POST`
- **Request Body:**
```
{
  "charge_id": "ch_1IuR7X2eZvKYlo2C8z0Bv2Cj"
}
```
- **Response:**
```
{
  "refund_id": "re_1IuS8H2eZvKYlo2C8z0Bv2Cf",
  "amount": 2000,
  "status": "succeeded"
}
```

## Docker Deployment

### Running with Docker

To containerize and run `rr-payments` as a Docker container:

1. **Build the Docker image:**
   ```sh
   docker build -t rr-payments .
   ```
2. **Run the container:**
   ```sh
   docker run -d --name rr-payments -p 8082:8082 --env-file=secrets.env rr-payments
   ```
3. **Verify the service is running:**
   ```sh
   curl http://localhost:8082/health
   ```

### Running with Docker Compose

To run `rr-payments` as part of a multi-container setup:

1. **Ensure Docker Compose is installed.**
2. **Start the service using Compose:**
   ```sh
   docker-compose -f docker-compose.yml up -d --build
   ```
3. **Check running containers:**
   ```sh
   docker ps
   ```
4. **Stop the service:**
   ```sh
   docker-compose down
   ```

### Ensuring Network Connectivity

If using a shared network between microservices:
```sh
docker network create shared_network || true
```
Ensure `rr-payments` is connected:
```sh
docker network connect shared_network rr-payments
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

- **Language:** Go
- **Framework:** Standard Library with `net/http`
- **Stripe SDK:** `github.com/stripe/stripe-go/v72`
- **Environment Management:** `github.com/joho/godotenv`
- **Containerization:** Docker, Docker Compose

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or collaboration, please reach out via LinkedIn or check out more of my work on [GitHub](https://www.github.com/tyler-pritchard/rr-payments).
