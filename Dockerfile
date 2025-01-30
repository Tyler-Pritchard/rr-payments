# Use an official Go image with the correct version
FROM golang:1.23.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o rr-payments ./cmd/payment-service/main.go

# Use a minimal Alpine image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Copy the built application from the builder stage
COPY --from=builder /app/rr-payments .

# Expose the port that the service runs on
EXPOSE 8082

# Run the application
CMD [ "./rr-payments" ]
