package controllers

import (
	"encoding/json"
	"glofox-backend/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// StudioController handles studio-related operations
type StudioController struct {
	DB *gorm.DB
}

// NewStudioController creates a new studio controller
func NewStudioController(db *gorm.DB) *StudioController {
	return &StudioController{DB: db}
}

// CreateStudio handles creation of a new studio
// @Summary Create a new studio
// @Description Create a new fitness studio
// @Tags studios
// @Accept json
// @Produce json
// @Param studio body models.Studio true "Studio information"
// @Success 201 {object} models.Studio
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /studios [post]
func (sc *StudioController) CreateStudio(w http.ResponseWriter, r *http.Request) {
	var studio models.Studio

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&studio); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Create studio in database
	if err := sc.DB.Create(&studio).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, studio)
}

// GetAllStudios retrieves all studios
// @Summary Get all studios
// @Description Get all fitness studios
// @Tags studios
// @Produce json
// @Success 200 {array} models.Studio
// @Failure 500 {object} ErrorResponse
// @Router /studios [get]
func (sc *StudioController) GetAllStudios(w http.ResponseWriter, r *http.Request) {
	var studios []models.Studio

	if err := sc.DB.Find(&studios).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, studios)
}

// GetStudio retrieves a specific studio by ID
// @Summary Get a studio by ID
// @Description Get a fitness studio by its ID
// @Tags studios
// @Produce json
// @Param id path int true "Studio ID"
// @Success 200 {object} models.Studio
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /studios/{id} [get]
func (sc *StudioController) GetStudio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid studio ID")
		return
	}

	var studio models.Studio
	if err := sc.DB.First(&studio, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Studio not found")
		return
	}

	respondWithJSON(w, http.StatusOK, studio)
}

// UpdateStudio updates a specific studio by ID
// @Summary Update a studio
// @Description Update a studio's information
// @Tags studios
// @Accept json
// @Produce json
// @Param id path int true "Studio ID"
// @Param studio body models.Studio true "Studio information"
// @Success 200 {object} models.Studio
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /studios/{id} [put]
func (sc *StudioController) UpdateStudio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid studio ID")
		return
	}

	var studio models.Studio
	if err := sc.DB.First(&studio, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Studio not found")
		return
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&studio); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Update studio in database
	if err := sc.DB.Save(&studio).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, studio)
}

// DeleteStudio deletes a specific studio by ID
// @Summary Delete a studio
// @Description Delete a studio by its ID
// @Tags studios
// @Produce json
// @Param id path int true "Studio ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /studios/{id} [delete]
func (sc *StudioController) DeleteStudio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid studio ID")
		return
	}

	var studio models.Studio
	if err := sc.DB.First(&studio, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Studio not found")
		return
	}

	if err := sc.DB.Delete(&studio).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Studio deleted successfully"})
}
