package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking represents a client booking for a class
type Booking struct {
	gorm.Model
	ClientName    string    `json:"client_name" gorm:"size:100;not null"`
	ClientEmail   string    `json:"client_email" gorm:"size:100;not null"`
	ClassID       uint      `json:"class_id"`
	Class         Class     `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	BookingStatus string    `json:"booking_status" gorm:"size:20;default:'confirmed'"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
