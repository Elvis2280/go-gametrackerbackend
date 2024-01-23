package models

import "gorm.io/gorm"

type Tags struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `gorm:"unique;not null" json:"name"`
}
