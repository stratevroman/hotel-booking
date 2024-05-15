package models

import (
	"applicationDesignTest/utils"
	"time"
)

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

type DecreaseQuotaDto struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
}

var Availability = []RoomAvailability{
	{"reddison", "lux", utils.Date(2024, 1, 1), 1},
	{"reddison", "lux", utils.Date(2024, 1, 2), 1},
	{"reddison", "lux", utils.Date(2024, 1, 3), 1},
	{"reddison", "lux", utils.Date(2024, 1, 4), 1},
	{"reddison", "lux", utils.Date(2024, 1, 5), 0},
}
