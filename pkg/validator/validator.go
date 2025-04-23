package validator

import (
	"errors"
	"time"
)

func ValidateDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date is required")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format. Use YYYY-MM-DD")
	}

	return date, nil
}

func ValidateDateRange(startDateStr, endDateStr string) (time.Time, time.Time, error) {
	startDate, err := ValidateDate(startDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endDate, err := ValidateDate(endDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if endDate.Before(startDate) {
		return time.Time{}, time.Time{}, errors.New("end date must be after start date")
	}

	return startDate, endDate, nil
}

func ValidateCapacity(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity must be a positive integer")
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ValidateClassName(className string) error {
	if className == "" {
		return errors.New("class name is required")
	}
	return nil
}

func ValidateBookingDate(dateStr string) (time.Time, error) {
	date, err := ValidateDate(dateStr)
	if err != nil {
		return time.Time{}, err
	}

	today := time.Now().Truncate(24 * time.Hour)
	if date.Before(today) {
		return time.Time{}, errors.New("cannot book a class for past dates")
	}

	return date, nil
}
