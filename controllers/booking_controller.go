package controllers

import (
	"encoding/json"
	"glofox-backend/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// BookingController handles booking-related operations
type BookingController struct {
	DB *gorm.DB
}

// NewBookingController creates a new booking controller
func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{DB: db}
}

// CreateBooking handles creation of a new booking
// @Summary Create a new booking
// @Description Create a new booking for a class
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.Booking true "Booking information"
// @Success 201 {object} models.Booking
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings [post]
func (bc *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	booking.BookingUUID = uuid.New().String()
	// Validate class exists
	var class models.Class
	if err := bc.DB.First(&class, booking.ClassID).Error; err != nil {
		respondWithError(w, http.StatusBadRequest, "Class not found")
		return
	}

	// Check if class has available capacity
	var bookingCount int64
	bc.DB.Model(&models.Booking{}).Where("class_id = ?", booking.ClassID).Count(&bookingCount)

	if int(bookingCount) >= class.Capacity {
		respondWithError(w, http.StatusBadRequest, "Class is fully booked")
		return
	}

	// Create booking in database
	if err := bc.DB.Create(&booking).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, booking)
}

// GetAllBookings retrieves all bookings
// @Summary Get all bookings
// @Description Get all bookings
// @Tags bookings
// @Produce json
// @Success 200 {array} models.Booking
// @Failure 500 {object} ErrorResponse
// @Router /bookings [get]
func (bc *BookingController) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking

	if err := bc.DB.Find(&bookings).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, bookings)
}

// GetBooking retrieves a specific booking by ID
// @Summary Get a booking by ID
// @Description Get a booking by its ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} models.Booking
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings/{id} [get]
func (bc *BookingController) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.First(&booking, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Booking not found")
		return
	}

	respondWithJSON(w, http.StatusOK, booking)
}

// GetBookingsByClass retrieves all bookings for a specific class
// @Summary Get bookings by class
// @Description Get all bookings for a specific class
// @Tags bookings
// @Produce json
// @Param classId path int true "Class ID"
// @Success 200 {array} models.Booking
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{classId}/bookings [get]
func (bc *BookingController) GetBookingsByClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classID, err := strconv.Atoi(vars["classId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	var bookings []models.Booking
	if err := bc.DB.Where("class_id = ?", classID).Find(&bookings).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, bookings)
}

// UpdateBooking updates a specific booking by ID
// @Summary Update a booking
// @Description Update a booking's information
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body models.Booking true "Booking information"
// @Success 200 {object} models.Booking
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings/{id} [put]
func (bc *BookingController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.First(&booking, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Booking not found")
		return
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Update booking in database
	if err := bc.DB.Save(&booking).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, booking)
}

// CancelBooking handles cancellation of an existing booking
// @Summary Cancel a booking
// @Description Cancel an existing booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} controllers.SuccessResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 404 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @Router /bookings/{id}/cancel [put]
func (bc *BookingController) CancelBooking(w http.ResponseWriter, r *http.Request) {
	// Extract booking ID from the request
	vars := mux.Vars(r)
	bookingID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	// Find the booking in the database
	var booking models.Booking
	if err := bc.DB.First(&booking, bookingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			respondWithError(w, http.StatusNotFound, "Booking not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Update the status to "cancelled"
	booking.Status = "cancelled"
	if err := bc.DB.Save(&booking).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with success
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Booking cancelled successfully"})
}

// DeleteBooking deletes a specific booking by ID
// @Summary Delete a booking
// @Description Delete a booking by its ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings/{id} [delete]
func (bc *BookingController) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid booking ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.First(&booking, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Booking not found")
		return
	}

	if err := bc.DB.Delete(&booking).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Booking deleted successfully"})
}
