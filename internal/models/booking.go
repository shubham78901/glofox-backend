// File: internal/models/booking.go

package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	ClassID   string    `json:"classId"`
	CreatedAt time.Time `json:"createdAt"`
}

type BookingInput struct {
	Name    string `json:"name" binding:"required"`
	Date    string `json:"date" binding:"required"`
	ClassID string `json:"classId" binding:"required"`
}

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

func NewBooking(input BookingInput) (*Booking, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	date, _ := time.Parse("2006-01-02", input.Date)

	return &Booking{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Date:      date,
		ClassID:   input.ClassID,
		CreatedAt: time.Now(),
	}, nil
}
