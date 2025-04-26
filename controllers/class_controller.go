package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"glofox-backend/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ClassController handles class-related operations
type ClassController struct {
	DB *gorm.DB
}

// NewClassController creates a new class controller
func NewClassController(db *gorm.DB) *ClassController {
	return &ClassController{DB: db}
}

// CreateClass handles creation of a new class
// @Summary Create a new class
// @Description Create a new fitness class
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.Class true "Class information"
// @Success 201 {object} models.ResClass
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes [post]
func (cc *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	var input models.Class

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input
	if input.StartTime.After(input.EndTime) {
		respondWithError(w, http.StatusBadRequest, "Start time must be before end time")
		return
	}

	// Create class from input
	class := models.Class{
		Name:        input.Name,
		Description: input.Description,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Capacity:    input.Capacity,
	}

	// Create class in database
	if err := cc.DB.Create(&class).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create class")
		return
	}

	// Prepare response
	resClass := models.ResClass{
		ClassUUID:   class.ClassUUID,
		CreatedAt:   class.CreatedAt,
		Name:        class.Name,
		Description: class.Description,
		StartTime:   class.StartTime,
		EndTime:     class.EndTime,
		Capacity:    class.Capacity,
	}

	respondWithJSON(w, http.StatusCreated, resClass)
}

// GetAllClasses retrieves all classes
// @Summary Get all classes
// @Description Get all fitness classes
// @Tags classes
// @Produce json
// @Param date query string false "Filter classes by date (YYYY-MM-DD)"
// @Success 200 {object} []models.ResClass
// @Failure 500 {object} ErrorResponse
// @Router /classes [get]
func (cc *ClassController) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	var classes []models.Class
	query := cc.DB

	// Add date filter if provided
	if date := r.URL.Query().Get("date"); date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
			return
		}
		query = query.Where("DATE(start_time) = ?", parsedDate.Format("2006-01-02"))
	}

	if err := query.Find(&classes).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve classes")
		return
	}

	// Convert to response models
	var resClasses []models.ResClass
	for _, class := range classes {
		resClasses = append(resClasses, models.ResClass{
			ClassUUID:   class.ClassUUID,
			CreatedAt:   class.CreatedAt,
			Name:        class.Name,
			Description: class.Description,
			StartTime:   class.StartTime,
			EndTime:     class.EndTime,
			Capacity:    class.Capacity,
		})
	}

	respondWithJSON(w, http.StatusOK, resClasses)
}

// GetClass retrieves a specific class by UUID
// @Summary Get a class by UUID
// @Description Get a fitness class by its UUID
// @Tags classes
// @Produce json
// @Param id path string true "Class UUID" Format(uuid)
// @Success 200 {object} models.ResClass
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{id} [get]
func (cc *ClassController) GetClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classUUID := vars["id"]

	// Validate UUID format
	if _, err := uuid.Parse(classUUID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	var class models.Class
	if err := cc.DB.Where("class_uuid = ?", classUUID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Class not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Database error")
		}
		return
	}

	resClass := models.ResClass{
		ClassUUID:   class.ClassUUID,
		CreatedAt:   class.CreatedAt,
		Name:        class.Name,
		Description: class.Description,
		StartTime:   class.StartTime,
		EndTime:     class.EndTime,
		Capacity:    class.Capacity,
	}

	respondWithJSON(w, http.StatusOK, resClass)
}
