# Build stage
FROM golang:1.22.6-alpine AS builder

WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .
RUN go build -o glofox-backend ./cmd/api

# Final stage
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/glofox-backend .

EXPOSE 8080
CMD ["./glofox-backend"]
