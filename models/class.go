package models

import (
	"time"
)

// Class represents a fitness class that can be booked
// @Description Fitness class information
type Class struct {
	// Standard gorm.Model fields
	ID        uint       `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2025-04-26T07:38:52Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-04-26T07:38:52Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`

	// Custom class fields
	Name        string    `json:"name" example:"Yoga Flow"`
	Description string    `json:"description" example:"A flowing yoga class for all levels"`
	StartTime   time.Time `json:"start_time" example:"2025-04-30T18:00:00Z"`
	EndTime     time.Time `json:"end_time" example:"2025-04-30T19:00:00Z"`
	Capacity    int       `json:"capacity" example:"20"`
	StudioID    uint      `json:"studio_id" example:"1"`
	Studio      Studio    `json:"studio" gorm:"foreignKey:StudioID" swaggerignore:"true"`
	Bookings    []Booking `json:"bookings,omitempty" gorm:"foreignKey:ClassID" swaggerignore:"true"`
}

// NewClass creates a new class
func NewClass(name, description string, startTime, endTime time.Time, capacity int, studioID uint) *Class {
	return &Class{
		Name:        name,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		Capacity:    capacity,
		StudioID:    studioID,
	}
}
