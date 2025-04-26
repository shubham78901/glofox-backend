
<img width="1470" alt="Screenshot 2025-04-23 at 9 28 36 PM" src="https://github.com/user-attachments/assets/41c7bc64-38fd-451b-9ae9-0d246c3815b4" />
<img width="1470" alt="Screenshot 2025-04-24 at 11 02 47 AM" src="https://github.com/user-attachments/assets/93cad8aa-2312-4099-a5bb-1d6c27b3306f" />
<img width="1470" alt="Screenshot 2025-04-24 at 11 01 49 AM" src="https://github.com/user-attachments/assets/2dbf65f2-baf1-4011-972a-2cb9cba12ead" />


<img width="1470" alt="Screenshot 2025-04-24 at 11 01 57 AM" src="https://github.com/user-attachments/assets/e654f86c-a526-455d-99b0-6bf93cc92475" />


# Glofox Studio API

A RESTful API for managing fitness studio classes and bookings, built with Go.

## Overview

This backend service provides endpoints for fitness studio management, allowing studio owners to:

- Create and manage recurring fitness classes
- Handle member bookings and attendance
- Track class capacity and availability
- Filter classes by date ranges

## Technology Stack

- **Go** - Core programming language
- **Gorilla Mux** - HTTP routing and request handling
- **Swagger/OpenAPI** - API documentation and specification
- **Testify** - Testing framework for assertions and mocks
- **GoMock** - Mocking framework for unit tests
- **UUID** - Unique identifier generation
- **Docker** - Containerization for deployment
- **Make** - Build automation

## Project Structure

```
glofox-backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── docs/                        # Swagger documentation
├── internal/
│   ├── api/
│   │   ├── handlers/            # HTTP request handlers
│   │   │   ├── booking.go       # Booking handler implementation
│   │   │   ├── booking_test.go  # Booking handler tests
│   │   │   ├── class.go         # Class handler implementation
│   │   │   └── class_test.go    # Class handler tests
│   │   ├── middleware/          # HTTP middleware
│   │   │   └── middleware.go    # Logger and error middleware
│   │   ├── responses/           # API response utilities
│   │   │   └── responses.go     # JSON response formatting
│   │   ├── router.go            # API route configuration
│   │   └── swagger.go           # Swagger setup
│   ├── models/                  # Domain models
│   │   ├── booking.go           # Booking model and validation
│   │   └── class.go             # Class model and validation
│   ├── mocks/                   # Auto-generated test mocks
│   │   ├── mock_booking_repository.go
│   │   └── mock_class_repository.go
│   └── repositories/            # Data access layer
│       ├── booking.go           # Booking repository implementation
│       └── class.go             # Class repository implementation
├── pkg/                         # Shared packages
├── Makefile                     # Build and deployment commands
├── Dockerfile                   # Docker container definition
└── README.md                    # Project documentation
```

## API Endpoints

### Classes

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/classes` | Create a new fitness class |
| `GET`  | `/classes` | Get all classes (with optional date filter) |
| `GET`  | `/classes/{id}` | Get a specific class by ID |

### Bookings

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/bookings` | Create a new booking |
| `GET`  | `/bookings` | Get all bookings |
| `GET`  | `/bookings/{id}` | Get a specific booking by ID |

## API Documentation

The API is documented using Swagger/OpenAPI. Once the application is running, you can access the documentation at:

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
# Clone the repository
git clone https://github.com/yourusername/glofox-backend.git
cd glofox-backend

# Set port (default is 8080)
export PORT=8080

# Run the application
go run cmd/api/main.go
```

## Testing

The application includes comprehensive unit tests for the handlers and models:

```bash
# Run all tests
go test ./...

# Run specific tests
go test ./internal/api/handlers -v
```

## Sample API Requests

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

### Get Classes by Date

```bash
curl -X GET "http://localhost:8080/classes?date=2023-05-15"
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

The application can be run in a Docker container. The Dockerfile provides a multi-stage build for optimized container size:

```bash
# Build Docker image
docker build -t glofox-api .

# Run in Docker container
docker run -p 8080:8080 glofox-api
```

Alternatively, you can use the provided Make commands:

```bash
# Build and run with Docker
make build
make run
```

## Design Decisions

- **In-memory Repository Pattern**: The application uses in-memory repositories for simplicity, but the interface design allows for easy substitution with a database implementation.
- **Thread-safe Operations**: Repository implementations use mutex locks to ensure thread safety for concurrent operations.
- **Validation**: Input validation is performed at the model level before data persistence.
- **Error Handling**: Consistent error responses are provided through the responses package.
- **Dependency Injection**: Handlers are initialized with their required repositories, making testing and future changes easier.
# Glofox Backend

A fitness class booking backend application built with Go, GORM, and Gorilla Mux.

## Features

- Manage fitness studios
- Create and manage classes
- Handle booking process
- RESTful API
- Swagger documentation
- Docker containerization

## Tech Stack

- **Language**: Go
- **Web Framework**: Gorilla Mux
- **ORM**: GORM
- **Database**: PostgreSQL
- **Documentation**: Swagger
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites

- Go (1.19+)
- Docker & Docker Compose
- PostgreSQL (if running locally)

### Running with Docker

1. Clone the repository
   ```sh
   git clone https://github.com/shubham78901/glofox-backend.git
   cd glofox-backend
   ```

2. Start the application using Docker Compose
   ```sh
   make docker-run
   ```

3. The API will be available at http://localhost:8080

### Running Locally

1. Clone the repository
   ```sh
   git clone https://github.com/shubham78901/glofox-backend.git
   cd glofox-backend
   ```

2. Install dependencies
   ```sh
   make deps
   ```

3. Set up environment variables (copy `.env.example` to `.env` and modify as needed)

4. Run the application
   ```sh
   make run
   ```

5. The API will be available at http://localhost:8080

## API Documentation

The API documentation is available at `/swagger/index.html` when the application is running.

### Main Endpoints

- **Studios**
  - `GET /api/v1/studios` - Get all studios
  - `POST /api/v1/studios` - Create a new studio
  - `GET /api/v1/studios/{id}` - Get a specific studio
  - `PUT /api/v1/studios/{id}` - Update a studio
  - `DELETE /api/v1/studios/{id}` - Delete a studio

- **Classes**
  - `GET /api/v1/classes` - Get all classes
  - `POST /api/v1/classes` - Create a new class
  - `GET /api/v1/classes/{id}` - Get a specific class
  - `PUT /api/v1/classes/{id}` - Update a class
  - `DELETE /api/v1/classes/{id}` - Delete a class
  - `GET /api/v1/studios/{studioId}/classes` - Get all classes for a studio

- **Bookings**
  - `GET /api/v1/bookings` - Get all bookings
  - `POST /api/v1/bookings` - Create a new booking
  - `GET /api/v1/bookings/{id}` - Get a specific booking
  - `PUT /api/v1/bookings/{id}` - Update a booking
  - `DELETE /api/v1/bookings/{id}` - Delete a booking
  - `PUT /api/v1/bookings/{id}/cancel` - Cancel a booking
  - `GET /api/v1/classes/{classId}/bookings` - Get all bookings for a class

## Development

### Useful Commands

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make fmt` - Format code
- `make swagger` - Generate Swagger documentation
- `make docker-build` - Build Docker image
- `make docker-up` - Start Docker containers
- `make docker-down` - Stop Docker containers
- `make all` - Clean, get dependencies, format code, build, test, and run Docker containers

## Project Structure

```
glofox-backend/
├── config/
│   └── database.go        # Database configuration
├── controllers/
│   ├── booking_controller.go
│   ├── class_controller.go
│   ├── studio_controller.go
│   └── utils.go
├── models/
│   ├── booking.go
│   ├── class.go
│   └── studio.go
├── routes/
│   ├── booking_routes.go
│   ├── class_routes.go
│   └── studio_routes.go
├── docs/                  # Swagger documentation
├── .env                   # Environment variables
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.