package models

import (
	"fmt"
	"time"
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o Order) String() string {
	return fmt.Sprintf("HotelID: %s, RoomID: %s, From: %s, To: %s", o.HotelID, o.RoomID, o.From, o.To)
}
