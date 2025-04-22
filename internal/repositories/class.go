// File: internal/repositories/class.go

package repositories

import (
	"errors"
	"glofox-backend/internal/models"
	"sync"
	"time"
)

// ClassRepository defines the interface for class storage operations
type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() []*models.Class
	GetByID(id string) (*models.Class, error)
	GetByDate(date time.Time) []*models.Class
}

// InMemoryClassRepository implements ClassRepository with in-memory storage
type InMemoryClassRepository struct {
	classes map[string]*models.Class
	mutex   sync.RWMutex
}

// NewClassRepository creates a new instance of InMemoryClassRepository
func NewClassRepository() ClassRepository {
	return &InMemoryClassRepository{
		classes: make(map[string]*models.Class),
	}
}

// Create adds a new class to the repository
func (r *InMemoryClassRepository) Create(class *models.Class) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.classes[class.ID] = class
	return nil
}

// GetAll returns all classes from the repository
func (r *InMemoryClassRepository) GetAll() []*models.Class {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	classes := make([]*models.Class, 0, len(r.classes))
	for _, class := range r.classes {
		classes = append(classes, class)
	}
	return classes
}

// GetByID returns a class by its ID
func (r *InMemoryClassRepository) GetByID(id string) (*models.Class, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	class, exists := r.classes[id]
	if !exists {
		return nil, errors.New("class not found")
	}
	return class, nil
}

// GetByDate returns all classes available on a given date
func (r *InMemoryClassRepository) GetByDate(date time.Time) []*models.Class {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	matchingClasses := make([]*models.Class, 0)
	for _, class := range r.classes {
		if class.IsDateInRange(date) {
			matchingClasses = append(matchingClasses, class)
		}
	}
	return matchingClasses
}
