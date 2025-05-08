package model

import "time"

type Experience struct {
	ID                uint      `gorm:"primaryKey"`
	Company           string    `json:"company"`
	Role              string    `json:"role"`
	StartMonth        string    `json:"start-month"`
	StartYear         string    `json:"start-year"`
	EndMonth          *string    `json:"end-month"`
	EndYear           *string    `json:"end-year"`
	IsCurrentlyWorking bool      `json:"currently-working"`
	Descriptions      []string `gorm:"type:text[]" json:"descriptions"`
	Logo              string    `json:"logo"`

	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
