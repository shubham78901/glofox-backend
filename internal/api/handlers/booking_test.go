package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"glofox-backend/internal/mocks"
	"glofox-backend/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)
	handler := NewBookingHandler(mockRepo)

	bookingInput := models.BookingInput{
		Name:    "John Doe",
		Date:    "2022-01-05",
		ClassID: "test-class-id",
	}
	requestBody, _ := json.Marshal(bookingInput)

	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	req := httptest.NewRequest("POST", "/bookings", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler.CreateBooking(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetBookingByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)
	handler := NewBookingHandler(mockRepo)

	mockBooking := &models.Booking{
		ID:        "test-id",
		Name:      "John Doe",
		Date:      time.Now(),
		ClassID:   "test-class-id",
		CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetByID("test-id").Return(mockBooking, nil)

	req := httptest.NewRequest("GET", "/bookings/test-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-id"})
	recorder := httptest.NewRecorder()

	handler.GetBookingByID(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetBookingByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)
	handler := NewBookingHandler(mockRepo)

	mockRepo.EXPECT().GetByID("non-existent-id").Return(nil, errors.New("booking not found"))

	req := httptest.NewRequest("GET", "/bookings/non-existent-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "non-existent-id"})
	recorder := httptest.NewRecorder()

	handler.GetBookingByID(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetAllBookings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)
	handler := NewBookingHandler(mockRepo)

	mockBookings := []*models.Booking{
		{ID: "test-id-1", Name: "John", Date: time.Now(), ClassID: "1", CreatedAt: time.Now()},
		{ID: "test-id-2", Name: "Jane", Date: time.Now(), ClassID: "2", CreatedAt: time.Now()},
	}

	mockRepo.EXPECT().GetAll().Return(mockBookings)

	req := httptest.NewRequest("GET", "/bookings", nil)
	recorder := httptest.NewRecorder()

	handler.GetAllBookings(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
