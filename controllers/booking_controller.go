package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"glofox-backend/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// BookingController handles booking-related operations
type BookingController struct {
	DB *gorm.DB
}

// NewBookingController creates a new booking controller instance
func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{DB: db}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a booking for a fitness class on a specific date
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.BookingInput true "Booking information"
// @Success 201 {object} models.BookingRes
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse "Class is full or not available on requested date"
// @Failure 500 {object} ErrorResponse
// @Router /bookings [post]
func (bc *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var input models.BookingInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input date
	bookingDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
		return
	}

	// Validate class exists and is available on the requested date
	var class models.Class
	if err := bc.DB.Where("class_uuid = ?", input.ClassID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Class not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Database error")
		}
		return
	}

	// Check if class is scheduled on the requested date

	if bookingDate.Before(class.StartTime) || bookingDate.After(class.EndTime) {
		respondWithError(w, http.StatusConflict, "Class not available on requested date")
		return
	}

	// Check capacity for the specific date
	var count int64
	if err := bc.DB.Model(&models.Booking{}).
		Where("class_id = ? AND date = ?", input.ClassID, input.Date).
		Count(&count).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to check capacity")
		return
	}

	if int(count) >= class.Capacity {
		respondWithError(w, http.StatusConflict, "Class is full for the requested date")
		return
	}

	// Create booking
	booking := models.Booking{
		ClassID:     input.ClassID,
		TraineeName: input.TraineeName,
		Date:        input.Date,
	}

	if err := bc.DB.Create(&booking).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create booking")
		return
	}

	// Convert to response model
	bookingRes := models.BookingRes{
		BookingUUID: booking.BookingUUID,
		CreatedAt:   booking.CreatedAt,
		ClassID:     booking.ClassID,
		TraineeName: booking.TraineeName,
		Date:        booking.Date,
	}

	respondWithJSON(w, http.StatusCreated, bookingRes)
}

// GetAllBookings godoc
// @Summary Get all bookings
// @Description Retrieve a list of all bookings with their details
// @Tags bookings
// @Produce json
// @Success 200 {array} models.BookingRes
// @Failure 500 {object} ErrorResponse
// @Router /bookings [get]
func (bc *BookingController) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking

	if err := bc.DB.Find(&bookings).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve bookings")
		return
	}

	// Convert to response models
	bookingsRes := make([]models.BookingRes, len(bookings))
	for i, b := range bookings {
		bookingsRes[i] = models.BookingRes{
			BookingUUID: b.BookingUUID,
			CreatedAt:   b.CreatedAt,
			ClassID:     b.ClassID,
			TraineeName: b.TraineeName,
			Date:        b.Date,
		}
	}

	respondWithJSON(w, http.StatusOK, bookingsRes)
}

// GetBooking godoc
// @Summary Get booking details
// @Description Get details of a specific booking by its UUID
// @Tags bookings
// @Produce json
// @Param id path string true "Booking UUID"
// @Success 200 {object} models.BookingRes
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings/{id} [get]
func (bc *BookingController) GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingUUID := vars["id"]

	var booking models.Booking
	if err := bc.DB.Where("booking_uuid = ?", bookingUUID).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Booking not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Database error")
		}
		return
	}

	bookingRes := models.BookingRes{
		BookingUUID: booking.BookingUUID,
		CreatedAt:   booking.CreatedAt,
		ClassID:     booking.ClassID,
		TraineeName: booking.TraineeName,
		Date:        booking.Date,
	}

	respondWithJSON(w, http.StatusOK, bookingRes)
}

// GetBookingsByClass godoc
// @Summary Get bookings by class
// @Description Retrieve all bookings for a specific class
// @Tags bookings
// @Produce json
// @Param classId path string true "Class UUID"
// @Success 200 {array} models.BookingRes
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{classId}/bookings [get]
func (bc *BookingController) GetBookingsByClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classID := vars["classId"]

	var bookings []models.Booking

	if err := bc.DB.Where("class_id = ?", classID).Find(&bookings).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve bookings")
		return
	}

	bookingsRes := make([]models.BookingRes, len(bookings))
	for i, b := range bookings {
		bookingsRes[i] = models.BookingRes{
			BookingUUID: b.BookingUUID,
			CreatedAt:   b.CreatedAt,
			ClassID:     b.ClassID,
			TraineeName: b.TraineeName,
			Date:        b.Date,
		}
	}

	respondWithJSON(w, http.StatusOK, bookingsRes)
}
