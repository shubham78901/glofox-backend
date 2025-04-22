// File: internal/api/handlers/booking.go

package handlers

import (
	"glofox-backend/internal/api/responses"
	"glofox-backend/internal/models"
	"glofox-backend/internal/repositories"

	"github.com/gin-gonic/gin"
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
// @Param booking body models.BookingInput true "Booking data"
// @Success 201 {object} responses.Response{data=models.Booking} "Booking created successfully"
// @Failure 400 {object} responses.Response "Invalid input or no class available for the date"
// @Failure 500 {object} responses.Response "Server error"
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var input models.BookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.BadRequestResponse(c, "Invalid input: "+err.Error())
		return
	}

	booking, err := models.NewBooking(input)
	if err != nil {
		responses.BadRequestResponse(c, err.Error())
		return
	}

	if err := h.repo.Create(booking); err != nil {
		responses.BadRequestResponse(c, err.Error())
		return
	}

	responses.CreatedResponse(c, "Booking created successfully", booking)
}

// GetAllBookings godoc
// @Summary Get all bookings
// @Description Retrieves a list of all bookings
// @Tags bookings
// @Produce json
// @Success 200 {object} responses.Response{data=[]models.Booking} "A list of bookings"
// @Failure 500 {object} responses.Response "Server error"
// @Router /bookings [get]
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings := h.repo.GetAll()
	responses.ListResponse(c, bookings, len(bookings))
}

// GetBookingByID godoc
// @Summary Get a booking by ID
// @Description Retrieves a booking by its ID
// @Tags bookings
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} responses.Response{data=models.Booking} "Booking details"
// @Failure 404 {object} responses.Response "Booking not found"
// @Failure 500 {object} responses.Response "Server error"
// @Router /bookings/{id} [get]
func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	id := c.Param("id")

	booking, err := h.repo.GetByID(id)
	if err != nil {
		responses.NotFoundResponse(c, "Booking not found")
		return
	}

	responses.OKResponse(c, booking)
}
