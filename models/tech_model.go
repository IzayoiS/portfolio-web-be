package model

import "time"

type TechStack struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
