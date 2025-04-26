package models

import (
	"time"

	"gorm.io/gorm"
)

// Class represents a fitness class
// @Description Class model for studio offerings
type Class struct {
	gorm.Model           // Embedded gorm.Model
	ID         uint      `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	StudioID   uint      `json:"studio_id" gorm:"not null"`
	Name       string    `json:"name" gorm:"not null"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time" gorm:"not null"`
	Capacity   int       `json:"capacity" gorm:"not null"`
	// Add swaggerignore to fields that cause recursion issues
	Studio    Studio     `json:"studio,omitempty" gorm:"foreignKey:StudioID" swaggerignore:"true"`
	Bookings  []Booking  `json:"bookings,omitempty" gorm:"foreignKey:ClassID" swaggerignore:"true"`
	CreatedAt time.Time  `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time  `json:"updated_at" swaggerignore:"true"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggerignore:"true"`
}
