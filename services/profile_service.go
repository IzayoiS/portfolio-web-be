package service

import (
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
)

func CreateProfile(profile *model.Profile) (*model.Profile, error) {
	if err := database.DB.Preload("User").Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func GetProfileByUserID(UserID uint) (*model.Profile, error) {
	var profile model.Profile
	if err := database.DB.Preload("User").Where("user_id = ?", UserID).First(&profile).Error; err != nil{
		return nil, err
	}
	return &profile, nil
}

func UpdateProfile(userID uint, updated *model.Profile) (*model.Profile, error) {
	var profile model.Profile

	if err := database.DB.Preload("User").Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	err := database.DB.Model(&profile).Updates(model.Profile{
		Name:         updated.Name,
		JobTitle:     updated.JobTitle,
		Bio:          updated.Bio,
		Location:     updated.Location,
		Availability: updated.Availability,
		ImageURL:     updated.ImageURL,
	}).Error

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
