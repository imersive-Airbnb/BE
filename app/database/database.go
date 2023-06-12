package database

import (
	"airnbn/app/config"
	rating "airnbn/features/rating/data"
	reservation "airnbn/features/reservation/data"
	room "airnbn/features/room/data"
	user "airnbn/features/user/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
func InitialMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&room.Room{},
		&rating.Rating{},
		&reservation.Reservation{},
	)
	if err != nil {
		return err
	}
	return nil
}
