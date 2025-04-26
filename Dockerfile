# Start from the official Go image with a newer version
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o glofox-backend .

# Start a new stage from scratch
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/glofox-backend .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./glofox-backend"]