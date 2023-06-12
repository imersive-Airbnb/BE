package data

import (
	rating "airnbn/features/rating/data"
	"time"

	reservation "airnbn/features/reservation/data"
)

type Room struct {
	RoomID      string `gorm:"primaryKey;type:varchar(50)"`
	Name        string `gorm:"type:varchar(100);not null;unique"`
	Price       uint
	Description string `gorm:"type:varchar(100);not null"`
	Latitude    uint
	Longitude   uint
	image       string
	CreatedAt   time.Time                 `gorm:"type:datetime"`
	UpdatedAt   time.Time                 `gorm:"type:datetime"`
	DeletedAt   time.Time                 `gorm:"type:datetime"`
	UserID      string                    `gorm:"type:varchar(50)"`
	User        User                      `gorm:"references:UserID"`
	Rating      []rating.Rating           `gorm:"foreignKey:RoomID"`
	reservation []reservation.Reservation `gorm:"foreignKey:RoomID"`
}

type User struct {
	UserID    string    `gorm:"primaryKey;type:varchar(50)"`
	name      string    `gorm:"type:varchar(100);not null;unique"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Status    string    `gorm:"type:enum('default', 'hosting'); default:'default'"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
	DeletedAt time.Time `gorm:"type:datetime"`
	Room      []Room    `gorm:"foreignKey:UserID"`
}
