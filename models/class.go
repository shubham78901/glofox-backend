// models/class.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ClassUUID   string         `json:"class_uuid" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	StartTime   time.Time      `json:"start_time"`
	EndTime     time.Time      `json:"end_time"`
	Capacity    int            `json:"capacity"`
	Bookings    []Booking      `json:"bookings" gorm:"foreignKey:ClassID;references:ClassUUID"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ClassUUID = uuid.New().String()
	return
}
