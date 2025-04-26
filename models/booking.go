package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking represents a booking for a class
// @Description Booking model for class reservations
type Booking struct {
	gorm.Model        // Embedded gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	ClassID    uint   `json:"class_id" gorm:"not null"`
	MemberName string `json:"member_name" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
	// Add swaggerignore to fields that cause recursion issues
	Class     Class      `json:"class,omitempty" gorm:"foreignKey:ClassID" swaggerignore:"true"`
	CreatedAt time.Time  `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time  `json:"updated_at" swaggerignore:"true"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggerignore:"true"`
}
