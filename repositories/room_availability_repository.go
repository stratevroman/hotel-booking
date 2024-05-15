package repositories

import (
	"applicationDesignTest/models"
	"errors"
	"time"
)

type RoomAvailabilityRepository interface {
	GetAvailabilityByDate(date time.Time) (*models.RoomAvailability, error)
	DecreaseQuotaByDate(date time.Time) error
}

type InMemoryRoomAvailabilityRepository struct {
	availabilities []models.RoomAvailability
}

func NewInMemoryRoomAvailabilityRepository(availabilities []models.RoomAvailability) *InMemoryRoomAvailabilityRepository {
	return &InMemoryRoomAvailabilityRepository{availabilities: availabilities}
}

func (r *InMemoryRoomAvailabilityRepository) GetAvailabilityByDate(date time.Time) (*models.RoomAvailability, error) {
	for _, availability := range r.availabilities {
		if availability.Date.Equal(date) {
			return &availability, nil
		}
	}

	return nil, nil
}

func (r *InMemoryRoomAvailabilityRepository) DecreaseQuotaByDate(date time.Time) error {
	for i, availability := range r.availabilities {
		if availability.Date.Equal(date) {
			if availability.Quota > 0 {
				r.availabilities[i].Quota--
				return nil
			} else {
				return errors.New("quota is already zero")
			}
		}
	}

	return errors.New("availability not found")
}
