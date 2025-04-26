package routes

import (
	"glofox-backend/controllers"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// SetupStudioRoutes configures the routes for studio operations
func SetupStudioRoutes(router *mux.Router, db *gorm.DB) {
	studioController := controllers.NewStudioController(db)

	// Studio routes
	router.HandleFunc("/studios", studioController.CreateStudio).Methods("POST")
	router.HandleFunc("/studios", studioController.GetAllStudios).Methods("GET")
	router.HandleFunc("/studios/{id}", studioController.GetStudio).Methods("GET")
	router.HandleFunc("/studios/{id}", studioController.UpdateStudio).Methods("PUT")
	router.HandleFunc("/studios/{id}", studioController.DeleteStudio).Methods("DELETE")
}
