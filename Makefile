# App Name
APP_NAME = glofox-backend

# Docker image tag
IMAGE_TAG = $(APP_NAME)

# Go build configuration
GO_BUILD_FLAGS = -v
MAIN_PATH = ./cmd/api

# API port
API_PORT = 8080
API_URL = http://localhost:$(API_PORT)

# Directories for tests
TEST_DIRS = ./internal/api/handlers

# Build Docker image
build:
	docker build -t $(IMAGE_TAG) .

# Run app in detached mode (remove previous container if exists)
run: stop build
	docker run -d -p $(API_PORT):$(API_PORT) --name $(APP_NAME) $(IMAGE_TAG)

# Stop and remove the container
stop:
	docker stop $(APP_NAME) 2>/dev/null || true
	docker rm $(APP_NAME) 2>/dev/null || true

# Start services with docker-compose
compose-up:
	docker-compose up -d

# Stop services
compose-down:
	docker-compose down

# Build the Go application locally
go-build:
	go build $(GO_BUILD_FLAGS) -o bin/$(APP_NAME) $(MAIN_PATH)

# Run the Go application locally
go-run:
	go run $(MAIN_PATH)

# Run unit tests
test:
	go test $(TEST_DIRS) -v

# Run a full test suite (unit tests and API tests)
test-all: test

.PHONY: build run stop compose-up compose-down go-build go-run test test-all
