// File: internal/models/booking.go

package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Booking represents a booking for a class
type Booking struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null;index"`
	ClassID   string    `json:"classId" gorm:"type:uuid;not null;index"`
	Class     *Class    `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for the Booking model
func (Booking) TableName() string {
	return "bookings"
}

// BookingInput represents the input for creating a new booking
type BookingInput struct {
	Name    string `json:"name" binding:"required"`
	Date    string `json:"date" binding:"required"`
	ClassID string `json:"classId" binding:"required"`
}

// Validate validates the booking input
func (bi *BookingInput) Validate() error {
	if bi.Name == "" {
		return errors.New("name is required")
	}

	if bi.ClassID == "" {
		return errors.New("classId is required")
	}

	_, err := time.Parse("2006-01-02", bi.Date)
	if err != nil {
		return errors.New("invalid date format. Use YYYY-MM-DD")
	}

	return nil
}

// NewBooking creates a new Booking instance from input
func NewBooking(input BookingInput) (*Booking, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	date, _ := time.Parse("2006-01-02", input.Date)

	return &Booking{
		ID:        generateUUID(),
		Name:      input.Name,
		Date:      date,
		ClassID:   input.ClassID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// generateUUID generates a new UUID string
func generateUUID() string {
	return uuid.New().String()
}
