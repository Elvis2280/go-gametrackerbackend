package models

import "gorm.io/gorm"

type Platforms struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	IconName   string `gorm:"not null" json:"iconName"`
}
