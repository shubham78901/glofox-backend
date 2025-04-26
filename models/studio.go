package models

import (
	"time"

	"gorm.io/gorm"
)

// Studio represents a fitness studio in the system
type Studio struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:100;not null"`
	Address     string    `json:"address" gorm:"size:255"`
	PhoneNumber string    `json:"phone_number" gorm:"size:20"`
	Email       string    `json:"email" gorm:"size:100;uniqueIndex"`
	TimeZone    string    `json:"time_zone" gorm:"size:50;default:'UTC'"`
	Classes     []Class   `json:"classes,omitempty" gorm:"foreignKey:StudioID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
