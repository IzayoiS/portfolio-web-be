package model

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string	`json:"-"`
	Email    string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Profile  *Profile `gorm:"foreignKey:UserID" json:"profile,omitempty"`
}