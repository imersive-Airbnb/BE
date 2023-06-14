package data

import (
	"airnbn/app/middlewares"
	"airnbn/features/user"
	"airnbn/helper"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// CheckProfileByID implements user.UserDataInterface.
func (repo *userQuery) CheckProfileByID(userID uuid.UUID) (user.UserCore, error) {
	var userData User

	// Find the user profile by ID
	tx := repo.db.First(&userData, "user_id = ?", userID)
	if tx.Error != nil {
		return user.UserCore{}, tx.Error
	}

	// Convert database model to user core model
	dataCore := user.UserCore{
		UserID:    userID.String(),
		Name:      userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return dataCore, nil
}

// UpgradeStatus implements user.UserDataInterface.
func (repo *userQuery) UpgradeStatus(userID string, newStatus string) error {
	// Dapatkan pengguna berdasarkan userID
	var user User
	if err := repo.db.First(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}

	// Validasi status baru
	if newStatus != "hosting" {
		return errors.New("invalid status")
	}

	// Perbarui status pengguna
	user.Status = newStatus

	// Simpan perubahan pada pengguna
	if err := repo.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUserByID implements user.UserDataInterface.
func (repo *userQuery) UpdateUserByID(userID string, updatedUser user.UserCore) error {
	// Dapatkan pengguna berdasarkan userID
	var user User
	if err := repo.db.First(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}

	// Perbarui bidang-bidang yang diperlukan pada user
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	// Jika password diisi, perbarui juga password
	if updatedUser.Password != "" {
		// Lakukan hashing password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	// Simpan perubahan pada pengguna
	if err := repo.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
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
		Phone:     user.Phone,
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
