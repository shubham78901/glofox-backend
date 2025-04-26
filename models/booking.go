// models/booking.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	BookingUUID string         `json:"booking_uuid" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ClassID     string         `json:"class_id"` // Foreign key to reference Class
	UserID      string         `json:"user_id"`
	Status      string         `json:"status"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.BookingUUID = uuid.New().String()
	return
}
