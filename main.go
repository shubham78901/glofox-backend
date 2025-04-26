package main

import (
	"glofox-backend/config"
	"glofox-backend/controllers"
	"glofox-backend/models"
	"glofox-backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

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
	r := mux.NewRouter()

	// Setup API routes
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	routes.SetupStudioRoutes(apiRouter, db)
	routes.SetupClassRoutes(apiRouter, db)
	routes.SetupBookingRoutes(apiRouter, db)

	// Add middleware
	r.Use(controllers.LoggingMiddleware)

	// Configure Swagger - specifically for /swagger/index.html
	// First, serve the /swagger/index.html endpoint
	r.HandleFunc("/swagger/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	// Then handle the general /swagger/ path
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Static files (if needed)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is healthy"))
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
