.PHONY: build run clean test docker-build docker-up docker-down swagger

# Run the application (stops previous containers, builds and runs in Docker)
run:
	docker-compose down
	docker-compose up --build

# Clean Docker resources
clean:
	docker-compose down -v
	docker system prune -f

# Run tests in a Docker container
test:
	docker-compose run --rm app go test -v ./...

# Format code in a Docker container
fmt:
	docker-compose run --rm app go fmt ./...

# Generate swagger documentation in a Docker container
swagger:
	docker-compose run --rm app swag init -g main.go -o ./docs

# Build Docker image
build:
	docker-compose build

# Run Docker containers in detached mode
up:
	docker-compose up -d

# Stop Docker containers
down:
	docker-compose down

# Show logs for running containers
logs:
	docker-compose logs -f

# Restart services
restart:
	docker-compose restart

# Show running containers
ps:
	docker-compose ps

# All-in-one command to restart fresh
all: clean build up