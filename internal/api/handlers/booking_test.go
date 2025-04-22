// File: internal/api/handlers/booking_test.go

package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"glofox-backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock BookingRepository
type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) Create(booking *models.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingRepository) GetAll() []*models.Booking {
	args := m.Called()
	return args.Get(0).([]*models.Booking)
}

func (m *MockBookingRepository) GetByID(id string) (*models.Booking, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Booking), args.Error(1)
}

func (m *MockBookingRepository) GetByClassAndDate(classID string, date time.Time) []*models.Booking {
	args := m.Called(classID, date)
	return args.Get(0).([]*models.Booking)
}

func TestCreateBooking(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	// Setup a valid booking input
	bookingInput := models.BookingInput{
		Name:    "John Doe",
		Date:    "2022-01-05",
		ClassID: "test-class-id",
	}
	requestBody, _ := json.Marshal(bookingInput)

	// Mock the repo.Create method to return nil error
	mockRepo.On("Create", mock.AnythingOfType("*models.Booking")).Return(nil)

	// Create a request
	req := httptest.NewRequest("POST", "/bookings", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.CreateBooking(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetBookingByID(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	// Create a mock booking to return
	mockBooking := &models.Booking{
		ID:        "test-id",
		Name:      "John Doe",
		Date:      time.Now(),
		ClassID:   "test-class-id",
		CreatedAt: time.Now(),
	}

	// Setup the mock
	mockRepo.On("GetByID", "test-id").Return(mockBooking, nil)

	// Create a request
	req := httptest.NewRequest("GET", "/bookings/test-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-id"})
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.GetBookingByID(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetBookingByID_NotFound(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	// Setup the mock to return an error
	mockRepo.On("GetByID", "non-existent-id").Return(nil, errors.New("booking not found"))

	// Create a request
	req := httptest.NewRequest("GET", "/bookings/non-existent-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "non-existent-id"})
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.GetBookingByID(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusNotFound, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetAllBookings(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	// Create mock bookings
	mockBookings := []*models.Booking{
		{
			ID:        "test-id-1",
			Name:      "John Doe",
			Date:      time.Now(),
			ClassID:   "test-class-id",
			CreatedAt: time.Now(),
		},
		{
			ID:        "test-id-2",
			Name:      "Jane Smith",
			Date:      time.Now(),
			ClassID:   "test-class-id",
			CreatedAt: time.Now(),
		},
	}

	// Setup the mock
	mockRepo.On("GetAll").Return(mockBookings)

	// Create a request
	req := httptest.NewRequest("GET", "/bookings", nil)
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.GetAllBookings(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}
