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

	bookingInput := models.BookingInput{
		Name:    "John Doe",
		Date:    "2022-01-05",
		ClassID: "test-class-id",
	}
	requestBody, _ := json.Marshal(bookingInput)

	mockRepo.On("Create", mock.AnythingOfType("*models.Booking")).Return(nil)

	req := httptest.NewRequest("POST", "/bookings", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler.CreateBooking(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	mockRepo.AssertExpectations(t)
}

func TestGetBookingByID(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	mockBooking := &models.Booking{
		ID:        "test-id",
		Name:      "John Doe",
		Date:      time.Now(),
		ClassID:   "test-class-id",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetByID", "test-id").Return(mockBooking, nil)

	req := httptest.NewRequest("GET", "/bookings/test-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-id"})
	recorder := httptest.NewRecorder()

	handler.GetBookingByID(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	mockRepo.AssertExpectations(t)
}

func TestGetBookingByID_NotFound(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

	mockRepo.On("GetByID", "non-existent-id").Return(nil, errors.New("booking not found"))

	req := httptest.NewRequest("GET", "/bookings/non-existent-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "non-existent-id"})
	recorder := httptest.NewRecorder()

	handler.GetBookingByID(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)

	mockRepo.AssertExpectations(t)
}

func TestGetAllBookings(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	handler := NewBookingHandler(mockRepo)

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

	mockRepo.On("GetAll").Return(mockBookings)

	req := httptest.NewRequest("GET", "/bookings", nil)
	recorder := httptest.NewRecorder()

	handler.GetAllBookings(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	mockRepo.AssertExpectations(t)
}
