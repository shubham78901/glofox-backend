.PHONY: run clean build up down migrate

# Run the application
run:
	docker-compose down -v
	docker-compose build
	docker-compose up

# Clean containers and networks
clean:
	docker-compose down

# Build containers
build:
	docker-compose build

# Start containers in background
up:
	docker-compose up -d

# Stop containers
down:
	docker-compose down

# Run migrations manually
migrate:
	docker-compose run --rm app migrate -path ./migrations -database $$DB_URL up