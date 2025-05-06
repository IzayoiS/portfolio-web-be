package service

import (
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
	if err := db.Where("Email = ? AND password = ?", email,password).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}