// File: internal/api/handlers/class.go

package handlers

import (
	"glofox-backend/internal/api/responses"
	"glofox-backend/internal/models"
	"glofox-backend/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
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
// @Description Creates a new class with the provided details
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.ClassInput true "Class data"
// @Success 201 {object} responses.Response{data=models.Class} "Class created successfully"
// @Failure 400 {object} responses.Response "Invalid input"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes [post]
func (h *ClassHandler) CreateClass(c *gin.Context) {
	var input models.ClassInput
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.BadRequestResponse(c, "Invalid input: "+err.Error())
		return
	}

	class, err := models.NewClass(input)
	if err != nil {
		responses.BadRequestResponse(c, err.Error())
		return
	}

	if err := h.repo.Create(class); err != nil {
		responses.InternalServerErrorResponse(c)
		return
	}

	responses.CreatedResponse(c, "Class created successfully", class)
}

// GetAllClasses godoc
// @Summary Get all classes
// @Description Retrieves a list of all classes
// @Tags classes
// @Produce json
// @Param date query string false "Optional date to filter classes (YYYY-MM-DD)"
// @Success 200 {object} responses.Response{data=[]models.Class} "A list of classes"
// @Failure 400 {object} responses.Response "Invalid date format"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes [get]
func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	dateParam := c.Query("date")
	if dateParam != "" {
		date, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			responses.BadRequestResponse(c, "Invalid date format. Use YYYY-MM-DD")
			return
		}

		classes := h.repo.GetByDate(date)
		responses.ListResponse(c, classes, len(classes))
		return
	}

	classes := h.repo.GetAll()
	responses.ListResponse(c, classes, len(classes))
}

// GetClassByID godoc
// @Summary Get a class by ID
// @Description Retrieves a class by its ID
// @Tags classes
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} responses.Response{data=models.Class} "Class details"
// @Failure 404 {object} responses.Response "Class not found"
// @Failure 500 {object} responses.Response "Server error"
// @Router /classes/{id} [get]
func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")

	class, err := h.repo.GetByID(id)
	if err != nil {
		responses.NotFoundResponse(c, "Class not found")
		return
	}

	responses.OKResponse(c, class)
}
