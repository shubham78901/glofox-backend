

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize repositories
	classRepo := repositories.NewClassRepository()
	bookingRepo := repositories.NewBookingRepository(classRepo)

	// Initialize handlers
	classHandler := handlers.NewClassHandler(classRepo)
	bookingHandler := handlers.NewBookingHandler(bookingRepo)

	// Setup router
	router := api.SetupRouter(classHandler, bookingHandler)

	// Start server
	serverAddr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
