package routes

import (
	"glofox-backend/controllers"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// SetupClassRoutes configures the routes for class operations
func SetupClassRoutes(router *mux.Router, db *gorm.DB) {
	classController := controllers.NewClassController(db)

	// Class routes
	router.HandleFunc("/classes", classController.CreateClass).Methods("POST")
	router.HandleFunc("/classes", classController.GetAllClasses).Methods("GET")
	router.HandleFunc("/classes/{id}", classController.GetClass).Methods("GET")

}
