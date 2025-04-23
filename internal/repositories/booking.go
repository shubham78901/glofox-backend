package repositories

import (
	"errors"
	"glofox-backend/internal/models"
	"sync"
	"time"
)

type BookingRepository interface {
	Create(booking *models.Booking) error
	GetAll() []*models.Booking
	GetByID(id string) (*models.Booking, error)
	GetByClassAndDate(classID string, date time.Time) []*models.Booking
}

type InMemoryBookingRepository struct {
	bookings  map[string]*models.Booking
	classRepo ClassRepository
	mutex     sync.RWMutex
}

func NewBookingRepository(classRepo ClassRepository) BookingRepository {
	return &InMemoryBookingRepository{
		bookings:  make(map[string]*models.Booking),
		classRepo: classRepo,
	}
}

func (r *InMemoryBookingRepository) Create(booking *models.Booking) error {
	// Check if class exists
	class, err := r.classRepo.GetByID(booking.ClassID)
	if err != nil {
		return errors.New("class not found")
	}

	// Check if booking date is within class date range
	if !class.IsDateInRange(booking.Date) {
		return errors.New("no class available on the requested date")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.bookings[booking.ID] = booking
	return nil
}

// GetAll returns all bookings from the repository
func (r *InMemoryBookingRepository) GetAll() []*models.Booking {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	bookings := make([]*models.Booking, 0, len(r.bookings))
	for _, booking := range r.bookings {
		bookings = append(bookings, booking)
	}
	return bookings
}

// GetByID returns a booking by its ID
func (r *InMemoryBookingRepository) GetByID(id string) (*models.Booking, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	booking, exists := r.bookings[id]
	if !exists {
		return nil, errors.New("booking not found")
	}
	return booking, nil
}

// GetByClassAndDate returns all bookings for a specific class on a specific date
func (r *InMemoryBookingRepository) GetByClassAndDate(classID string, date time.Time) []*models.Booking {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	normalizedDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	matchingBookings := make([]*models.Booking, 0)

	for _, booking := range r.bookings {
		bookingDate := time.Date(booking.Date.Year(), booking.Date.Month(), booking.Date.Day(), 0, 0, 0, 0, time.UTC)
		if booking.ClassID == classID && bookingDate.Equal(normalizedDate) {
			matchingBookings = append(matchingBookings, booking)
		}
	}

	return matchingBookings
}
