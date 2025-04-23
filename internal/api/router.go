package api

import (
	"encoding/json"
	"net/http"

	"glofox-backend/internal/api/handlers"
	"glofox-backend/internal/api/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(classHandler *handlers.ClassHandler, bookingHandler *handlers.BookingHandler) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.ErrorHandler)

	SetupSwagger(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message":       "Welcome to Glofox Studio API",
			"documentation": "/swagger/index.html",
		})
	}).Methods("GET")

	router.HandleFunc("/classes", classHandler.CreateClass).Methods("POST")
	router.HandleFunc("/classes", classHandler.GetAllClasses).Methods("GET")
	router.HandleFunc("/classes/{id}", classHandler.GetClassByID).Methods("GET")

	router.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings", bookingHandler.GetAllBookings).Methods("GET")
	router.HandleFunc("/bookings/{id}", bookingHandler.GetBookingByID).Methods("GET")

	return router
}
