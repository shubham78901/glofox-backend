package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Booking represents a trainee's reservation for a class on a specific date
// @Description Booking information
type Booking struct {
	BookingUUID string         `json:"-" gorm:"primaryKey" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt   time.Time      `json:"created_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
	ClassID     string         `json:"class_id" gorm:"size:36;not null" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	TraineeName string         `json:"trainee_name" gorm:"size:100;not null" example:"John Doe"`
	Date        string         `json:"date" gorm:"type:date;not null" example:"2025-04-26"`
}

// BookingRes represents the response format for booking data
// @Description Booking response information
type BookingRes struct {
	BookingUUID string    `json:"booking_uuid" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt   time.Time `json:"created_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	ClassID     string    `json:"class_id" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	TraineeName string    `json:"trainee_name" example:"John Doe"`
	Date        string    `json:"date" example:"2025-04-26"`
}

// BookingInput represents the input format for creating a booking
// @Description Booking creation input
type BookingInput struct {
	ClassID     string `json:"class_id" binding:"required" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	TraineeName string `json:"trainee_name" binding:"required" example:"John Doe"`
	Date        string `json:"date" binding:"required" example:"2025-04-26"`
}

func (c *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	c.BookingUUID = uuid.NewString()
	return nil
}
