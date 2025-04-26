package models

import (
	"time"

	"gorm.io/gorm"
)

// Studio represents a fitness studio
// @Description Studio model for fitness businesses
type Studio struct {
	gorm.Model        // Embedded gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name       string `json:"name" gorm:"not null"`
	Address    string `json:"address" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
	Phone      string `json:"phone" gorm:"not null"`
	// Add swaggerignore to fields that cause recursion issues
	Classes   []Class    `json:"classes,omitempty" gorm:"foreignKey:StudioID" swaggerignore:"true"`
	CreatedAt time.Time  `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time  `json:"updated_at" swaggerignore:"true"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggerignore:"true"`
}
