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
// @Success 201 {object} models.Class
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes [post]
func (cc *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	class.ClassUUID = uuid.New().String()

	// Create class in database
	if err := cc.DB.Create(&class).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, class)
}

// GetAllClasses retrieves all classes
// @Summary Get all classes
// @Description Get all fitness classes
// @Tags classes
// @Produce json
// @Success 200 {array} models.Class
// @Failure 500 {object} ErrorResponse
// @Router /classes [get]
func (cc *ClassController) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	var classes []models.Class

	if err := cc.DB.Find(&classes).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, classes)
}

// GetClass retrieves a specific class by ID
// @Summary Get a class by ID
// @Description Get a fitness class by its ID
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} models.Class
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{id} [get]
func (cc *ClassController) GetClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	var class models.Class
	if err := cc.DB.First(&class, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Class not found")
		return
	}

	respondWithJSON(w, http.StatusOK, class)
}

// UpdateClass updates a specific class by ID
// @Summary Update a class
// @Description Update a class's information
// @Tags classes
// @Accept json
// @Produce json
// @Param id path int true "Class ID"
// @Param class body models.Class true "Class information"
// @Success 200 {object} models.Class
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{id} [put]
func (cc *ClassController) UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	var class models.Class
	if err := cc.DB.First(&class, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Class not found")
		return
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Update class in database
	if err := cc.DB.Save(&class).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, class)
}

// DeleteClass deletes a specific class by ID
// @Summary Delete a class
// @Description Delete a class by its ID
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classes/{id} [delete]
func (cc *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	var class models.Class
	if err := cc.DB.First(&class, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Class not found")
		return
	}

	if err := cc.DB.Delete(&class).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Class deleted successfully"})
}
