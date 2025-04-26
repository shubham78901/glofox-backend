package models

import (
	"time"
)

// Studio represents a fitness studio that offers classes
// @Description Fitness studio information
type Studio struct {
	// Standard gorm.Model fields
	ID        uint       `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2025-04-26T07:38:52Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-04-26T07:38:52Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`

	// Custom studio fields
	Name        string  `json:"name" example:"Glofox Fitness Studio"`
	Description string  `json:"description" example:"A premium fitness studio in downtown"`
	Address     string  `json:"address" example:"123 Main Street, City"`
	Phone       string  `json:"phone" example:"+1-123-456-7890"`
	Email       string  `json:"email" example:"info@glofoxstudio.com"`
	Classes     []Class `json:"classes,omitempty" gorm:"foreignKey:StudioID" swaggerignore:"true"`
}

// NewStudio creates a new studio
func NewStudio(name, description, address, phone, email string) *Studio {
	return &Studio{
		Name:        name,
		Description: description,
		Address:     address,
		Phone:       phone,
		Email:       email,
	}
}
