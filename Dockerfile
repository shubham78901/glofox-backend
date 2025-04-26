# Build stage
FROM golang:1.22.6-alpine AS builder
WORKDIR /app

# Install dependencies
RUN apk add --no-cache gcc musl-dev

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .
RUN go build -o glofox-backend ./cmd/api

# Final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/glofox-backend .

EXPOSE 8080

# Set up environment variables (defaults, will be overridden by docker-compose)
ENV PORT=8080 \
    DB_HOST=postgres \
    DB_PORT=5432 \
    DB_USER=glofox \
    DB_PASSWORD=glofox123 \
    DB_NAME=glofox \
    DB_SSLMODE=disable

CMD ["./glofox-backend"]