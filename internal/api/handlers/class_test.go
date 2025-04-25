package handlers

import (
	"bytes"
	"encoding/json"
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

func TestCreateClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockClassRepository(ctrl)
	handler := NewClassHandler(mockRepo)

	classInput := models.ClassInput{
		ClassName: "Test Class",
		StartDate: "2022-01-01",
		EndDate:   "2022-01-10",
		Capacity:  10,
	}
	requestBody, _ := json.Marshal(classInput)

	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	req := httptest.NewRequest("POST", "/classes", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler.CreateClass(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestGetClassByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockClassRepository(ctrl)
	handler := NewClassHandler(mockRepo)

	mockClass := &models.Class{
		ID:        "test-id",
		ClassName: "Test Class",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 10),
		Capacity:  10,
		CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetByID("test-id").Return(mockClass, nil)

	req := httptest.NewRequest("GET", "/classes/test-id", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-id"})
	recorder := httptest.NewRecorder()

	handler.GetClassByID(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllClasses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockClassRepository(ctrl)
	handler := NewClassHandler(mockRepo)

	mockClasses := []*models.Class{
		{ID: "test-id-1", ClassName: "Class 1", StartDate: time.Now(), EndDate: time.Now(), Capacity: 10, CreatedAt: time.Now()},
		{ID: "test-id-2", ClassName: "Class 2", StartDate: time.Now(), EndDate: time.Now(), Capacity: 12, CreatedAt: time.Now()},
	}

	mockRepo.EXPECT().GetAll().Return(mockClasses)

	req := httptest.NewRequest("GET", "/classes", nil)
	recorder := httptest.NewRecorder()

	handler.GetAllClasses(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
