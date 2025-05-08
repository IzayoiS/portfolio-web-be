package service

import (
	"errors"
	"log"
	"os"
	model "portfolio-web-be/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func LoginUser(email, password string) (*model.User, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var user model.User

	if err := db.Preload("Profile").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found with the given email")
		}
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
