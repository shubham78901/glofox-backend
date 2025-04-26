package models

import (
	"time"
)

// Booking represents a user booking for a class
// @Description Booking information for classes
type Booking struct {
	// Standard gorm.Model fields
	ID        uint       `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2025-04-26T07:38:52Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-04-26T07:38:52Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`

	// Custom booking fields
	UserID   uint      `json:"user_id" example:"1"`
	ClassID  uint      `json:"class_id" example:"2"`
	BookedAt time.Time `json:"booked_at" example:"2025-04-26T07:38:52Z"`
	Status   string    `json:"status" gorm:"default:'confirmed'" example:"confirmed"`
	Class    Class     `json:"class" gorm:"foreignKey:ClassID" swaggerignore:"true"`
}

// NewBooking creates a new booking
func NewBooking(userID, classID uint) *Booking {
	return &Booking{
		UserID:   userID,
		ClassID:  classID,
		BookedAt: time.Now(),
		Status:   "confirmed", // Default status
	}
}
