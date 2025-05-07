package service

import (
	"log"
	"os"
	model "portfolio-web-be/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func AuthUser(email, password string) (*model.User, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return nil,err
	}

	var user model.User

	if err := db.Preload("Profile").Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User not found with the given credentials")
		}
		return nil, err
	}
	return &user, nil
}