package repositories

import (
	"applicationDesignTest/models"
	"errors"
	"sync"
	"time"
)

type RoomAvailabilityRepository interface {
	GetAvailability(date time.Time, hotelId string, roomId string) (*models.RoomAvailability, error)
	DecreaseQuotaByDate(date time.Time, hotelId string, roomId string) error
}

type InMemoryRoomAvailabilityRepository struct {
	mutex          *sync.Mutex
	availabilities []models.RoomAvailability
}

func NewInMemoryRoomAvailabilityRepository(mutex *sync.Mutex, availabilities []models.RoomAvailability) *InMemoryRoomAvailabilityRepository {
	return &InMemoryRoomAvailabilityRepository{mutex: mutex, availabilities: availabilities}
}

func (r *InMemoryRoomAvailabilityRepository) GetAvailability(date time.Time, hotelId string, roomId string) (*models.RoomAvailability, error) {
	availability := findAvailabilityIndex(r.availabilities, date, hotelId, roomId)
	if availability == -1 {
		return nil, nil
	}

	return &r.availabilities[availability], nil
}

func (r *InMemoryRoomAvailabilityRepository) DecreaseQuotaByDate(date time.Time, hotelId string, roomId string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	availability := findAvailabilityIndex(r.availabilities, date, hotelId, roomId)
	if availability == -1 {
		return errors.New("availability not found")
	}

	if r.availabilities[availability].Quota > 0 {
		r.availabilities[availability].Quota--
		return nil
	}

	return errors.New("quota is already zero")
}

func findAvailabilityIndex(availabilities []models.RoomAvailability, date time.Time, hotelId string, roomId string) int {
	for i, availability := range availabilities {
		if availability.Date.Equal(date) && availability.HotelID == hotelId && availability.RoomID == roomId {
			return i
		}
	}
	return -1
}
