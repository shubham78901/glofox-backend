// File: internal/models/class.go

package models

import (
	"errors"
	"time"
)

// Class represents a fitness class in the system
type Class struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	ClassName string    `json:"className" gorm:"not null"`
	StartDate time.Time `json:"startDate" gorm:"not null"`
	EndDate   time.Time `json:"endDate" gorm:"not null"`
	Capacity  int       `json:"capacity" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for the Class model
func (Class) TableName() string {
	return "classes"
}

// ClassInput represents the input for creating a new class
type ClassInput struct {
	ClassName string `json:"className" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required,min=1"`
}

// Validate validates the class input
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

// NewClass creates a new Class instance from input
func NewClass(input ClassInput) (*Class, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	startDate, _ := time.Parse("2006-01-02", input.StartDate)
	endDate, _ := time.Parse("2006-01-02", input.EndDate)

	return &Class{
		ID:        generateUUID(),
		ClassName: input.ClassName,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  input.Capacity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// IsDateInRange checks if a date is within the class's date range
func (c *Class) IsDateInRange(date time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	startDate := time.Date(c.StartDate.Year(), c.StartDate.Month(), c.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(c.EndDate.Year(), c.EndDate.Month(), c.EndDate.Day(), 0, 0, 0, 0, time.UTC)

	return (date.Equal(startDate) || date.After(startDate)) && (date.Equal(endDate) || date.Before(endDate))
}
