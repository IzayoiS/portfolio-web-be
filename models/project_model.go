package model

import (
	"time"
)

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProjectName string	  `json:"project_name"`
	Descriptions string   `json:"descriptions"`
	TechStack   string	  `json:"tech_stack"`
	LinkGithub  string	  `json:"link_github"`
	LinkWebsite string	  `json:"link_website"`
	Logo        string	  `json:"logo"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
