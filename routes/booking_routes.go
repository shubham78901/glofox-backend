package routes

import (
	"glofox-backend/controllers"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// SetupBookingRoutes configures the routes for booking operations
func SetupBookingRoutes(router *mux.Router, db *gorm.DB) {
	bookingController := controllers.NewBookingController(db)

	// Booking routes
	router.HandleFunc("/bookings", bookingController.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings", bookingController.GetAllBookings).Methods("GET")
	router.HandleFunc("/bookings/{id}", bookingController.GetBooking).Methods("GET")
	router.HandleFunc("/bookings/{id}", bookingController.UpdateBooking).Methods("PUT")
	router.HandleFunc("/bookings/{id}", bookingController.DeleteBooking).Methods("DELETE")
	router.HandleFunc("/bookings/{id}/cancel", bookingController.CancelBooking).Methods("PUT")

	// Class-specific booking routes
	router.HandleFunc("/classes/{classId}/bookings", bookingController.GetBookingsByClass).Methods("GET")
}
