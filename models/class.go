package models

import (
	"time"

	"gorm.io/gorm"
)

// Class represents a fitness class in the system
type Class struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"size:500"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	Capacity    int       `json:"capacity" gorm:"default:20"`
	StudioID    uint      `json:"studio_id"`
	Studio      Studio    `json:"studio,omitempty" gorm:"foreignKey:StudioID"`
	Bookings    []Booking `json:"bookings,omitempty" gorm:"foreignKey:ClassID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
