// File: cmd/api/main.go

package main

import (
	"log"

	_ "glofox-backend/docs"
	"glofox-backend/internal/api"
	"glofox-backend/internal/api/handlers"
	"glofox-backend/internal/repositories"
)

// @title           Glofox Studio API
// @version         1.0
// @description     API for managing studio classes and bookings
// @host            localhost:8080
// @BasePath        /
func main() {
	// Initialize repositories
	classRepo := repositories.NewClassRepository()
	bookingRepo := repositories.NewBookingRepository(classRepo)

	// Initialize handlers
	classHandler := handlers.NewClassHandler(classRepo)
	bookingHandler := handlers.NewBookingHandler(bookingRepo)

	// Setup router
	router := api.SetupRouter(classHandler, bookingHandler)

	log.Println("Server is running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
