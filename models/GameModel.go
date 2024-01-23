package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Platforms   []Platforms `gorm:"many2many:game_platforms;"`
	Tags        []Tags      `gorm:"many2many:game_tags;"`
}
