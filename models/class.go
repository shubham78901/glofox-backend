package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Class represents a fitness class that can be booked
// @Description Class information (database model)
type Class struct {
	ClassUUID   string         `json:"-" gorm:"primaryKey" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt   time.Time      `json:"-" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	UpdatedAt   time.Time      `json:"-" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
	Name        string         `json:"name" example:"Yoga Class"`
	Description string         `json:"description" example:"A relaxing yoga session for beginners"`
	StartTime   time.Time      `json:"start_time" swaggertype:"string" format:"date-time" example:"2025-04-26T09:00:00Z"`
	EndTime     time.Time      `json:"end_time" swaggertype:"string" format:"date-time" example:"2025-04-26T10:00:00Z"`
	Capacity    int            `json:"capacity" example:"20" minimum:"1"`
	Bookings    []Booking      `json:"-" gorm:"foreignKey:ClassID;references:ClassUUID" swaggerignore:"true"`
}

// ResClass represents the response format for Class data
// @Description Class information in API responses
type ResClass struct {
	ClassUUID   string    `json:"classId" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt   time.Time `json:"created_at" swaggertype:"string" format:"date-time" example:"2025-04-26T00:00:00Z"`
	Name        string    `json:"name" example:"Yoga Class"`
	Description string    `json:"description" example:"A relaxing yoga session for beginners"`
	StartTime   time.Time `json:"start_time" swaggertype:"string" format:"date-time" example:"2025-04-26T09:00:00Z"`
	EndTime     time.Time `json:"end_time" swaggertype:"string" format:"date-time" example:"2025-04-26T10:00:00Z"`
	Capacity    int       `json:"capacity" example:"20" minimum:"1"`
}

// BeforeCreate generates a new UUID before creating a new Class record
func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ClassUUID = uuid.NewString()
	return nil
}
