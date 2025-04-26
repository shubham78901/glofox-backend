// models/booking.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking represents a user's reservation for a class
// @Description Booking information
type Booking struct {
	BookingUUID string         `json:"-" gorm:"primaryKey" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt   time.Time      `json:"created_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
	ClassID     string         `json:"class_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID      string         `json:"user_id" example:"user123"`
	Status      string         `json:"status" example:"confirmed"`
}
