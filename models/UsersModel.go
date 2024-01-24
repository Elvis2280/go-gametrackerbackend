package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"not null" json:"username"`
	Password   string `gorm:"not null" json:"password"`
	Email      string `gorm:"unique; not null" json:"email"`
	IsVerified bool   `gorm:"default:false" json:"isVerified"`
	Games      []Game `gorm:"foreignKey:UserID" json:"games"`
}

type UserLogin struct {
	Email    string `gorm:"unique; not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserSignup struct {
	UserLogin
	Username string `gorm:"not null" json:"username"`
}
