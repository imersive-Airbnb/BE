package data

import (
	"time"
)

type Reservation struct {
	ReservationID  string    `gorm:"primaryKey;type:varchar(50)"`
	Name           string    `gorm:"type:varchar(100);not null;unique"`
	Description    string    `gorm:"type:varchar(100);not null"`
	start_date     time.Time `gorm:"type:date"`
	end_date       time.Time `gorm:"type:date"`
	Price          uint
	Qty            uint
	Total_price    uint
	Status         string `gorm:"type:varchar(100)"`
	Payment_type   string
	Payment_method string
	CreatedAt      time.Time `gorm:"type:datetime"`
	UpdatedAt      time.Time `gorm:"type:datetime"`
	DeletedAt      time.Time `gorm:"type:datetime"`
	UserID         string    `gorm:"type:varchar(50)"`
	RoomID         string    `gorm:"type:varchar(50)"`
	User           User      `gorm:"references:UserID"`
	Room           Room      `gorm:"references:RoomID"`
}
type Room struct {
	RoomID      string `gorm:"primaryKey;type:varchar(50)"`
	Name        string `gorm:"type:varchar(100);not null;unique"`
	Price       uint
	Description string `gorm:"type:varchar(100);not null"`
	Latitude    float32
	Longitude   float32
	image       string
	CreatedAt   time.Time     `gorm:"type:datetime"`
	UpdatedAt   time.Time     `gorm:"type:datetime"`
	DeletedAt   time.Time     `gorm:"type:datetime"`
	UserID      string        `gorm:"type:varchar(50)"`
	User        User          `gorm:"references:UserID"`
	Reservation []Reservation `gorm:"foreignKey:RoomID"`
}

type User struct {
	UserID      string        `gorm:"primaryKey;type:varchar(50)"`
	name        string        `gorm:"type:varchar(100);not null;unique"`
	Email       string        `gorm:"type:varchar(100);not null;unique"`
	Password    string        `gorm:"type:varchar(100);not null"`
	Status      string        `gorm:"type:enum('default', 'hosting'); default:'default'"`
	CreatedAt   time.Time     `gorm:"type:datetime"`
	UpdatedAt   time.Time     `gorm:"type:datetime"`
	DeletedAt   time.Time     `gorm:"type:datetime"`
	Room        []Room        `gorm:"foreignKey:UserID"`
	Reservation []Reservation `gorm:"foreignKey:UserID"`
}
