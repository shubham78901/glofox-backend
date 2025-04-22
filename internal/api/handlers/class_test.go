// File: internal/api/handlers/class_test.go

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"glofox-backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock ClassRepository
type MockClassRepository struct {
	mock.Mock
}

func (m *MockClassRepository) Create(class *models.Class) error {
	args := m.Called(class)
	return args.Error(0)
}

func (m *MockClassRepository) GetAll() []*models.Class {
	args := m.Called()
	return args.Get(0).([]*models.Class)
}

func (m *MockClassRepository) GetByID(id string) (*models.Class, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Class), args.Error(1)
}

func (m *MockClassRepository) GetByDate(date time.Time) []*models.Class {
	args := m.Called(date)
	return args.Get(0).([]*models.Class)
}

func TestCreateClass(t *testing.T) {
	mockRepo := new(MockClassRepository)
	handler := NewClassHandler(mockRepo)

	// Setup a valid class input
	classInput := models.ClassInput{
		ClassName: "Test Class",
		StartDate: "2022-01-01",
		EndDate:   "2022-01-10",
		Capacity:  10,
	}
	requestBody, _ := json.Marshal(classInput)

	// Mock the repo.Create method to return nil error
	mockRepo.On("Create", mock.AnythingOfType("*models.Class")).Return(nil)

	// Create a request
	req := httptest.NewRequest("POST", "/classes", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.CreateClass(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetClassByID(t *testing.T) {
	mockRepo := new(MockClassRepository)
	handler := NewClassHandler(mockRepo)

	// Create a mock class to return
	mockClass := &models.Class{
		ID:        "test-id",
		ClassName: "Test Class",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 10),
		Capacity:  10,
		CreatedAt: time.Now(),
	}

	// Setup the mock
	mockRepo.On("GetByID", "test-id").Return(mockClass, nil)

	// Create a request
	req := httptest.NewRequest("GET", "/classes/test-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-id"})
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.GetClassByID(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetAllClasses(t *testing.T) {
	mockRepo := new(MockClassRepository)
	handler := NewClassHandler(mockRepo)

	// Create mock classes
	mockClasses := []*models.Class{
		{
			ID:        "test-id-1",
			ClassName: "Test Class 1",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 10),
			Capacity:  10,
			CreatedAt: time.Now(),
		},
		{
			ID:        "test-id-2",
			ClassName: "Test Class 2",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 10),
			Capacity:  15,
			CreatedAt: time.Now(),
		},
	}

	// Setup the mock
	mockRepo.On("GetAll").Return(mockClasses)

	// Create a request
	req := httptest.NewRequest("GET", "/classes", nil)
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.GetAllClasses(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}
