// File: internal/repositories/booking.go

package repositories

import (
	"errors"
	"glofox-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// BookingRepository defines the interface for booking operations
type BookingRepository interface {
	Create(booking *models.Booking) error
	GetAll() ([]*models.Booking, error)
	GetByID(id string) (*models.Booking, error)
	GetByClassAndDate(classID string, date time.Time) ([]*models.Booking, error)
}

// DBBookingRepository implements BookingRepository using GORM
type DBBookingRepository struct {
	db *gorm.DB
}

// NewBookingRepository creates a new booking repository
func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &DBBookingRepository{db: db}
}

// Create inserts a new booking into the database
func (r *DBBookingRepository) Create(booking *models.Booking) error {
	// Check if class exists
	var class models.Class
	if err := r.db.Where("id = ?", booking.ClassID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("class not found")
		}
		return err
	}

	// Check if booking date is within class date range
	if !class.IsDateInRange(booking.Date) {
		return errors.New("no class available on the requested date")
	}

	// Check if class has reached capacity
	var bookingCount int64
	formattedDate := booking.Date.Format("2006-01-02")
	if err := r.db.Model(&models.Booking{}).
		Where("class_id = ? AND date = ?", booking.ClassID, formattedDate).
		Count(&bookingCount).Error; err != nil {
		return err
	}

	if int(bookingCount) >= class.Capacity {
		return errors.New("class is fully booked for this date")
	}

	// Create the booking
	return r.db.Create(booking).Error
}

// GetAll retrieves all bookings from the database
func (r *DBBookingRepository) GetAll() ([]*models.Booking, error) {
	var bookings []*models.Booking
	if err := r.db.Preload("Class").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

// GetByID retrieves a booking by its ID
func (r *DBBookingRepository) GetByID(id string) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.Preload("Class").Where("id = ?", id).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &booking, nil
}

// GetByClassAndDate retrieves bookings for a specific class on a specific date
func (r *DBBookingRepository) GetByClassAndDate(classID string, date time.Time) ([]*models.Booking, error) {
	var bookings []*models.Booking
	formattedDate := date.Format("2006-01-02")

	if err := r.db.Preload("Class").
		Where("class_id = ? AND date = ?", classID, formattedDate).
		Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}
