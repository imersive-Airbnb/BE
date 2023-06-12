package service

import (
	"airnbn/features/user"
	"airnbn/helper"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (user.UserCore, string, error) {
	if email == "" || password == "" {
		return user.UserCore{}, "", errors.New("email dan password harus diisi")
	}

	// Mencari pengguna berdasarkan email
	user, token, err := service.userData.Login(email, password)
	if err != nil {
		return user, "", err
	}

	return user, token, nil
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(user user.UserCore) error {
	// Validasi input pengguna
	err := service.validate.Struct(user)
	if err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(user.Password) // Use the HashPassword function from the helper package
	if err != nil {
		return err
	}

	// Update password dengan hashed password
	user.Password = hashedPassword

	// Insert pengguna ke dalam database
	err = service.userData.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
