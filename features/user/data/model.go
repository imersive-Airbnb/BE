package data

import (
	rating "airnbn/features/rating/data"
	reservation "airnbn/features/reservation/data"
	room "airnbn/features/room/data"
	"airnbn/features/user"
	"time"
)

type User struct {
	UserID      string                    `gorm:"primaryKey;type:varchar(50)"`
	Name        string                    `gorm:"type:varchar(100);not null;unique"`
	Phone       string                    `gorm:"type:varchar(100);not null"`
	Email       string                    `gorm:"type:varchar(100);not null;unique"`
	Password    string                    `gorm:"type:varchar(100);not null"`
	Status      string                    `gorm:"type:enum('default', 'hosting'); default:'default'"`
	CreatedAt   time.Time                 `gorm:"type:datetime"`
	UpdatedAt   time.Time                 `gorm:"type:datetime"`
	DeletedAt   time.Time                 `gorm:"type:datetime;default:null"`
	Room        []room.Room               `gorm:"foreignKey:UserID"`
	Rating      []rating.Rating           `gorm:"foreignKey:UserID"`
	reservation []reservation.Reservation `gorm:"foreignKey:UserID"`
}

func userModels(u User) user.UserCore {
	return user.UserCore{
		UserID:    u.UserID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Password:  u.Password,
		Status:    u.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
