package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Image       string      `json:"image"`
	Platforms   []Platforms `gorm:"many2many:game_platforms;"`
	Tags        []Tags      `gorm:"many2many:game_tags;"`
	Email       string      `json:"email" gorm:"type:text"`
}

type CreateGame struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Image       string      `json:"image" binding:"required"`
	Status      string      `json:"status" binding:"required"`
	Platforms   []Platforms `json:"platforms" binding:"required"`
	Tags        []Tags      `json:"tags" binding:"required"`
	UserID      string      `json:"user_id" binding:"required"`
}
