package user

import (
	"airnbn/features/room"
	"time"

	"github.com/google/uuid"
)

type UserCore struct {
	UserID    string
	Name      string
	Email     string
	Phone     string
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
	CheckProfileByID(userID uuid.UUID) (UserCore, error)
	// UpdateUserByID(userID uuid.UUID, updatedUser UserCore) error
	// UpgradeStatus(userID uuid.UUID, newStatus string) error
	UpdateUserByID(userID string, updatedUser UserCore) error
	UpgradeStatus(userID string, newStatus string) error
}

type UserServiceInterface interface {
	Create(user UserCore) error
	Login(email, password string) (UserCore, string, error)
	CheckProfile(userID uuid.UUID) (UserCore, error)
	// UpdateUser(userID uuid.UUID, updatedUser UserCore) error
	// UpgradeStatus(userID uuid.UUID, newStatus string) error
	UpdateUser(userID string, updatedUser UserCore) error
	UpgradeStatus(userID string, newStatus string) error
}
