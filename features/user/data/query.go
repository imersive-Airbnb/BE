package data

import (
	"airnbn/app/middlewares"
	"airnbn/features/user"
	"airnbn/helper"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (user.UserCore, string, error) {
	var userData User

	// Mencocokkan data inputan email dengan email di database
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return user.UserCore{}, "", errors.New("login failed, email salah")
		}
		return user.UserCore{}, "", tx.Error
	}

	// Mencocokkan data inputan password dengan password yang telah dihash di database
	checkPassword := helper.CheckPasswordHash(userData.Password, password)
	if !checkPassword {
		return user.UserCore{}, "", errors.New("login failed, password salah")
	}

	dataCore := user.UserCore{
		UserID:    userData.UserID, // Convert UserID to string
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Status:    userData.Status,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	token, err := middlewares.CreateToken(uuid.ClockSequence())
	if err != nil {
		return user.UserCore{}, "", err
	}

	return dataCore, token, nil
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(user user.UserCore) error {
	model := User{
		UserID:    uuid.New().String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return repo.db.Create(&model).Error
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
