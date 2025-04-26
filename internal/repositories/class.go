// File: internal/repositories/class.go

package repositories

import (
	"errors"
	"glofox-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// ClassRepository defines the interface for class operations
type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() ([]*models.Class, error)
	GetByID(id string) (*models.Class, error)
	GetByDate(date time.Time) ([]*models.Class, error)
}

// DBClassRepository implements ClassRepository using GORM
type DBClassRepository struct {
	db *gorm.DB
}

// NewClassRepository creates a new class repository
func NewClassRepository(db *gorm.DB) ClassRepository {
	return &DBClassRepository{db: db}
}

// Create inserts a new class into the database
func (r *DBClassRepository) Create(class *models.Class) error {
	return r.db.Create(class).Error
}

// GetAll retrieves all classes from the database
func (r *DBClassRepository) GetAll() ([]*models.Class, error) {
	var classes []*models.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

// GetByID retrieves a class by its ID
func (r *DBClassRepository) GetByID(id string) (*models.Class, error) {
	var class models.Class
	if err := r.db.Where("id = ?", id).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("class not found")
		}
		return nil, err
	}
	return &class, nil
}

// GetByDate retrieves classes available on a specific date
func (r *DBClassRepository) GetByDate(date time.Time) ([]*models.Class, error) {
	var classes []*models.Class
	formattedDate := date.Format("2006-01-02")

	if err := r.db.Where("start_date <= ? AND end_date >= ?", formattedDate, formattedDate).Find(&classes).Error; err != nil {
		return nil, err
	}

	return classes, nil
}
