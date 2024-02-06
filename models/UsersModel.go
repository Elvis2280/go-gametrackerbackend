package models

type User struct {
	Username   string `gorm:"not null" json:"username"`
	Password   string `gorm:"not null" json:"password"`
	Email      string `gorm:"unique; not null; primarykey" json:"email"`
	IsVerified bool   `gorm:"default:false" json:"isVerified"`
	Games      []Game `gorm:"foreignKey:Email" json:"games"`
}

type UserLogin struct {
	Email    string `gorm:"unique; not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserSignup struct {
	UserLogin
	Username string `gorm:"not null" json:"username"`
}
