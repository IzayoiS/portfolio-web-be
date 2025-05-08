package model

import (
	"time"

	"github.com/lib/pq"
)

type Project struct {
	ID          uint      `gorm:"primaryKey"`
	ProjectName string	  `json:"project-name"`
	Descriptions pq.StringArray `gorm:"type:text[]" json:"descriptions"`
	TechStack   string	  `json:"tech-stack"`
	LinkGithub  string	  `json:"link-github"`
	LinkWebsite string	  `json:"link-website"`
	Logo        string	  `json:"logo"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
