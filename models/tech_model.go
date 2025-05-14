package model

import "time"

type TechStack struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"uniqueIndex:idx_user_name" json:"name"`
	Icon   string `json:"icon"`

	UserID uint   `gorm:"uniqueIndex:idx_user_name"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
