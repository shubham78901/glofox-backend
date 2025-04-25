// File: internal/api/handlers/booking.go
// mockgen -source=internal/repositories/class.go -destination=internal/mocks/mock_class_repository.go -package=mocks
package handlers

import (
	"encoding/json"
	"net/http"

	"glofox-backend/internal/api/responses"
	"glofox-backend/internal/models"
	"glofox-backend/internal/repositories"

	"github.com/gorilla/mux"
)

// BookingHandler handles HTTP requests related to bookings
type BookingHandler struct {
	repo repositories.BookingRepository
}

// NewBookingHandler creates a new BookingHandler instance
func NewBookingHandler(repo repositories.BookingRepository) *BookingHandler {
	return &BookingHandler{repo: repo}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Creates a new booking for a member to attend a class
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.BookingInput true "Booking information"
// @Success 201 {object} responses.Response{data=models.Booking} "Booking created successfully"
// @Failure 400 {object} responses.Response "Invalid input"
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var input models.BookingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.BadRequestResponse(w, "Invalid input: "+err.Error())
		return
	}

	booking, err := models.NewBooking(input)
	if err != nil {
		responses.BadRequestResponse(w, err.Error())
		return
	}

	if err := h.repo.Create(booking); err != nil {
		responses.BadRequestResponse(w, err.Error())
		return
	}

	responses.CreatedResponse(w, "Booking created successfully", booking)
}

// GetAllBookings godoc
// @Summary Get all bookings
// @Description Retrieves a list of all bookings
// @Tags bookings
// @Produce json
// @Success 200 {object} responses.Response{data=[]models.Booking} "List of bookings"
// @Router /bookings [get]
func (h *BookingHandler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := h.repo.GetAll()
	responses.ListResponse(w, bookings, len(bookings))
}

// GetBookingByID godoc
// @Summary Get booking by ID
// @Description Retrieves a booking by its ID
// @Tags bookings
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} responses.Response{data=models.Booking} "Booking found"
// @Failure 404 {object} responses.Response "Booking not found"
// @Router /bookings/{id} [get]
func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	booking, err := h.repo.GetByID(id)
	if err != nil {
		responses.NotFoundResponse(w, "Booking not found")
		return
	}

	responses.OKResponse(w, booking)
}
