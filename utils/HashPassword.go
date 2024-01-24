package utils

import (
	"gametracker/db"
	"gametracker/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost of the hashing algorithm

	if err != nil { // if there is an error, return empty string and error
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password string, email string) (bool, error) {
	database := db.GetDatabase()
	var user models.User

	if err := database.Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}
