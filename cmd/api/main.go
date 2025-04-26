// File: cmd/api/main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "glofox-backend/docs"
	"glofox-backend/internal/api"
	"glofox-backend/internal/api/handlers"
	"glofox-backend/internal/database"
	"glofox-backend/internal/models"
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

	// Setup database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database schema
	if err := db.AutoMigrate(&models.Class{}, &models.Booking{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	classRepo := repositories.NewClassRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)

	// Initialize handlers
	classHandler := handlers.NewClassHandler(classRepo)
	bookingHandler := handlers.NewBookingHandler(bookingRepo)

	// Setup router
	router := api.SetupRouter(classHandler, bookingHandler)

	// Setup server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	log.Printf("Server is running on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
