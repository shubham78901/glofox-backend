// File: internal/api/handlers/class.go

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"glofox-backend/internal/api/responses"
	"glofox-backend/internal/models"
	"glofox-backend/internal/repositories"

	"github.com/gorilla/mux"
)

// ClassHandler handles HTTP requests related to classes
type ClassHandler struct {
	repo repositories.ClassRepository
}

// NewClassHandler creates a new ClassHandler instance
func NewClassHandler(repo repositories.ClassRepository) *ClassHandler {
	return &ClassHandler{repo: repo}
}

// CreateClass godoc
// @Summary Create a new class
// @Description Creates a new fitness class with the provided details
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.ClassInput true "Class information"
// @Success 201 {object} responses.Response{data=models.Class} "Class created successfully"
// @Failure 400 {object} responses.Response "Invalid input"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes [post]
func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var input models.ClassInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.BadRequestResponse(w, "Invalid input: "+err.Error())
		return
	}

	class, err := models.NewClass(input)
	if err != nil {
		responses.BadRequestResponse(w, err.Error())
		return
	}

	if err := h.repo.Create(class); err != nil {
		responses.InternalServerErrorResponse(w)
		return
	}

	responses.CreatedResponse(w, "Class created successfully", class)
}

// GetAllClasses godoc
// @Summary Get all classes
// @Description Retrieves a list of all classes, optionally filtered by date
// @Tags classes
// @Produce json
// @Param date query string false "Filter classes by date (YYYY-MM-DD)"
// @Success 200 {object} responses.Response{data=[]models.Class} "List of classes"
// @Failure 400 {object} responses.Response "Invalid date format"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes [get]
func (h *ClassHandler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	dateParam := r.URL.Query().Get("date")
	if dateParam != "" {
		date, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			responses.BadRequestResponse(w, "Invalid date format. Use YYYY-MM-DD")
			return
		}

		classes, err := h.repo.GetByDate(date)
		if err != nil {
			responses.InternalServerErrorResponse(w)
			return
		}

		responses.ListResponse(w, classes, len(classes))
		return
	}

	classes, err := h.repo.GetAll()
	if err != nil {
		responses.InternalServerErrorResponse(w)
		return
	}

	responses.ListResponse(w, classes, len(classes))
}

// GetClassByID godoc
// @Summary Get class by ID
// @Description Retrieves a class by its ID
// @Tags classes
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} responses.Response{data=models.Class} "Class found"
// @Failure 404 {object} responses.Response "Class not found"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes/{id} [get]
func (h *ClassHandler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	class, err := h.repo.GetByID(id)
	if err != nil {
		responses.NotFoundResponse(w, "Class not found")
		return
	}

	responses.OKResponse(w, class)
}
