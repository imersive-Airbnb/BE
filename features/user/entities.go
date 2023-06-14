package user

import (
	"airnbn/features/room"
	"time"
)

type UserCore struct {
	UserID    string
	Name      string
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Room      []room.RoomCore
}

type UserDataInterface interface {
	Insert(user UserCore) error
	Login(email, password string) (UserCore, string, error)
	CheckProfileByID(userID string) (UserCore, error)
	UpdateUserByID(userID string, updatedUser UserCore) error
	UpgradeStatus(userID string, newStatus string) error
}

type UserServiceInterface interface {
	Create(user UserCore) error
	Login(email, password string) (UserCore, string, error)
	CheckProfile(userID string) (UserCore, error)
	UpdateUser(userID string, updatedUser UserCore) error
	UpgradeStatus(userID string, newStatus string) error
}
