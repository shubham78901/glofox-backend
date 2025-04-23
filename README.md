
<img width="1470" alt="Screenshot 2025-04-23 at 9 28 36 PM" src="https://github.com/user-attachments/assets/41c7bc64-38fd-451b-9ae9-0d246c3815b4" />
<img width="1470" alt="Screenshot 2025-04-23 at 9 27 03 PM" src="https://github.com/user-attachments/assets/619a9cfe-678c-419f-815f-92fadc33e028" />
<img width="1470" alt="Screenshot 2025-04-23 at 9 27 03 PM" src="https://github.com/user-attachments/assets/f08b791c-f734-4eba-af56-5d50f21f65b2" />
<img width="1470" alt="Screenshot 2025-04-23 at 9 29 02 PM" src="https://github.com/user-attachments/assets/8a65cd40-c9ee-4653-b363-a9ee03da6b60" />




# Glofox Studio API

A backend API for managing fitness studio classes and bookings.

## Overview

This project provides a RESTful API for fitness studio management, allowing:
- Creation and management of fitness classes
- Member booking management
- Class schedule exploration

## Technology Stack

- **Go** - Core programming language
- **Gorilla Mux** - HTTP router and URL matcher
- **Swagger/OpenAPI** - API documentation
- **testify** - Testing framework
- **JSON** - Data interchange format
- **Docker** - Containerization
- **Make** - Build automation

## Project Structure

```
glofox-backend/
  ├── cmd/
  │   └── api/
  │       └── main.go            # Application entry point
  ├── docs/                      # Swagger documentation
  ├── internal/
  │   ├── api/
  │   │   ├── handlers/          # HTTP request handlers
  │   │   ├── middleware/        # HTTP middleware
  │   │   ├── responses/         # API response utilities
  │   │   └── router.go          # API routes configuration
  │   ├── models/                # Domain models + Validation utilities
  │   ├── repositories/          # Data access layer
  │ 
  ├── Makefile                   # Build and deployment commands
  ├── Dockerfile                 # Docker container definition
  └── README.md                  # This file
```

## API Endpoints

### Classes

- `POST /classes` - Create a new fitness class
- `GET /classes` - Get all classes (with optional date filter)
- `GET /classes/{id}` - Get a specific class by ID

### Bookings

- `POST /bookings` - Create a new booking
- `GET /bookings` - Get all bookings
- `GET /bookings/{id}` - Get a specific booking by ID

## API Documentation

The API is documented using Swagger/OpenAPI 2.0. Access the documentation at:

```
http://localhost:8080/swagger/index.html
```

## Running the Application

### Using Make

The project includes a Makefile with various commands to simplify development and deployment:

```bash
# Build Docker image
make build

# Run application in Docker
make run

# Stop running Docker container
make stop

# Start all services using docker-compose
make compose-up

# Stop all services
make compose-down

# Build the Go application locally
make go-build

# Run the application locally
make go-run

# Run unit tests
make test

# Run all tests (unit and API tests)
make test-all
```

### Manual Setup

```bash
# Set port (default is 8080)
export PORT=8080

# Run the application
go run cmd/api/main.go
```

## Testing

```bash

# Run specific tests
go test ./internal/api/handlers -v
```

## Sample Requests

### Create a Class

```bash
curl -X POST http://localhost:8080/classes \
  -H "Content-Type: application/json" \
  -d '{
    "className": "Yoga Basics",
    "startDate": "2023-05-01",
    "endDate": "2023-05-31",
    "capacity": 15
  }'
```

### Create a Booking

```bash
curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "date": "2023-05-15",
    "classId": "class-id-here"
  }'
```

## Docker Support

The application can be run in a Docker container. The Makefile provides convenient commands for building and running the Docker image:

```bash
# Build Docker image
make build

# Run in Docker container
make run
```

Alternatively, you can use docker-compose:

```bash
# Start all services
make compose-up
```

## License

This project is licensed under the MIT License.
