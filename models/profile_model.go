package model

import "time"

type Profile struct {
	ID 			 uint 	`gorm:"primaryKey" json:"id"`
	UserID       uint   `gorm:"uniqueIndex" json:"user_id"`
	Name         string `json:"name"`
	JobTitle     string `json:"job_title"`
	Bio          string `json:"bio"`
	Location     string `json:"location"`
	Availability string `json:"availability"`
	ImageURL     string `json:"image_url"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time

	User *User `gorm:"foreignKey:UserID" json:"User"`
}