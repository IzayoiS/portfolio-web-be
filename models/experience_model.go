package model

import (
	"time"

	"github.com/lib/pq"
)

type Experience struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Company           string    `json:"company"`
	Role              string    `json:"role"`
	StartMonth        string    `json:"start_month"`
	StartYear         string    `json:"start_year"`
	EndMonth          *string    `json:"end_month"`
	EndYear           *string    `json:"end_year"`
	IsCurrentlyWorking *bool      `json:"currently_working"`
	Descriptions      pq.StringArray `gorm:"type:text[]" json:"descriptions"`
	Logo              string    `json:"logo"`

	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
