package room

import (
	"time"
)

type RoomCore struct {
	RoomID      string
	Name        string
	Price       uint
	Description string
	Latitude    uint
	Longitude   uint
	image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	UserID      string
}
