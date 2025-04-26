package main

import (
	"glofox-backend/config"
	"glofox-backend/controllers"
	"glofox-backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors" // Added missing import

	httpSwagger "github.com/swaggo/http-swagger"
	// Make sure this import path matches your project structure
	_ "glofox-backend/docs"
)

// @title Glofox API
// @version 1.0
// @description Glofox Backend API service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@glofox.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment")
	}

	// Initialize database
	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database models
	db.AutoMigrate(&models.Class{}, &models.Studio{}, &models.Booking{})

	// Initialize router
	router := mux.NewRouter()

	// Create API subrouter
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize controllers
	studioController := controllers.NewStudioController(db)
	classController := controllers.NewClassController(db)
	bookingController := controllers.NewBookingController(db)

	// Register routes
	// Studio routes
	apiRouter.HandleFunc("/studios", studioController.CreateStudio).Methods("POST")
	apiRouter.HandleFunc("/studios", studioController.GetAllStudios).Methods("GET")
	apiRouter.HandleFunc("/studios/{id}", studioController.GetStudio).Methods("GET")
	apiRouter.HandleFunc("/studios/{id}", studioController.UpdateStudio).Methods("PUT")
	apiRouter.HandleFunc("/studios/{id}", studioController.DeleteStudio).Methods("DELETE")

	// Class routes
	apiRouter.HandleFunc("/classes", classController.CreateClass).Methods("POST")
	apiRouter.HandleFunc("/classes", classController.GetAllClasses).Methods("GET")
	apiRouter.HandleFunc("/classes/{id}", classController.GetClass).Methods("GET")
	apiRouter.HandleFunc("/classes/{id}", classController.UpdateClass).Methods("PUT")
	apiRouter.HandleFunc("/classes/{id}", classController.DeleteClass).Methods("DELETE")

	// Booking routes
	apiRouter.HandleFunc("/bookings", bookingController.CreateBooking).Methods("POST")
	apiRouter.HandleFunc("/bookings", bookingController.GetAllBookings).Methods("GET")
	apiRouter.HandleFunc("/bookings/{id}", bookingController.GetBooking).Methods("GET")
	apiRouter.HandleFunc("/bookings/{id}", bookingController.UpdateBooking).Methods("PUT")
	apiRouter.HandleFunc("/bookings/{id}", bookingController.DeleteBooking).Methods("DELETE")
	apiRouter.HandleFunc("/bookings/{id}/cancel", bookingController.CancelBooking).Methods("PUT")

	// Swagger documentation - using more explicit configuration
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	// Use CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	// Start server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
