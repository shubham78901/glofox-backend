// File: internal/models/class.go

package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Class represents a class offered by a studio
type Class struct {
	ID        string    `json:"id"`
	ClassName string    `json:"className"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"createdAt"`
}

// ClassInput represents the input data for creating a class
type ClassInput struct {
	ClassName string `json:"className" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required,min=1"`
}

// Validate validates the class input data
func (ci *ClassInput) Validate() error {
	if ci.ClassName == "" {
		return errors.New("className is required")
	}

	startDate, err := time.Parse("2006-01-02", ci.StartDate)
	if err != nil {
		return errors.New("invalid startDate format. Use YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", ci.EndDate)
	if err != nil {
		return errors.New("invalid endDate format. Use YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return errors.New("endDate must be after startDate")
	}

	if ci.Capacity < 1 {
		return errors.New("capacity must be at least 1")
	}

	return nil
}

// NewClass creates a new Class instance from ClassInput
func NewClass(input ClassInput) (*Class, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	startDate, _ := time.Parse("2006-01-02", input.StartDate)
	endDate, _ := time.Parse("2006-01-02", input.EndDate)

	return &Class{
		ID:        uuid.New().String(),
		ClassName: input.ClassName,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  input.Capacity,
		CreatedAt: time.Now(),
	}, nil
}

// IsDateInRange checks if a given date falls within the class's date range
func (c *Class) IsDateInRange(date time.Time) bool {
	// Normalize dates by setting time to midnight
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	startDate := time.Date(c.StartDate.Year(), c.StartDate.Month(), c.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(c.EndDate.Year(), c.EndDate.Month(), c.EndDate.Day(), 0, 0, 0, 0, time.UTC)

	return (date.Equal(startDate) || date.After(startDate)) && (date.Equal(endDate) || date.Before(endDate))
}
