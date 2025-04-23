package repositories

import (
	"errors"
	"glofox-backend/internal/models"
	"sync"
	"time"
)

type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() []*models.Class
	GetByID(id string) (*models.Class, error)
	GetByDate(date time.Time) []*models.Class
}

type InMemoryClassRepository struct {
	classes map[string]*models.Class
	mutex   sync.RWMutex
}

func NewClassRepository() ClassRepository {
	return &InMemoryClassRepository{
		classes: make(map[string]*models.Class),
	}
}

func (r *InMemoryClassRepository) Create(class *models.Class) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.classes[class.ID] = class
	return nil
}

func (r *InMemoryClassRepository) GetAll() []*models.Class {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	classes := make([]*models.Class, 0, len(r.classes))
	for _, class := range r.classes {
		classes = append(classes, class)
	}
	return classes
}

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
