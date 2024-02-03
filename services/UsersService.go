package services

import (
	"gametracker/db"
	"gametracker/models"
	"gametracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// SignUp godoc
// @Tags Authentication
// @Summary Create an account
// @ID create-account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /signup [POST]
// @Param platform body models.UserSignup true "Tag"
func SignUp(c *gin.Context) {
	database := db.GetDatabase()

	var user models.User

	parse := utils.CheckParse(c, &user) // check if the JSON is parsed correctly
	if parse == nil {
		return
	}

	user.Email = strings.ToLower(user.Email) // convert email to lowercase

	checkIfUserExist := database.Where("email = ?", strings.ToLower(user.Email)).First(&user)
	if checkIfUserExist.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	hash, err := utils.HashPassword(user.Password) // hash the password
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	user.Password = hash

	if err := database.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

// Login godoc
// @Tags Authentication
// @Summary Login into an account
// @ID login-account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /login [POST]
// @Param platform body models.UserLogin true "Tag"
func Login(c *gin.Context) {
	var user models.User
	database := db.GetDatabase()

	parse := utils.CheckParse(c, &user) // check if the JSON is parsed correctly
	if parse == nil {
		return
	}

	_, err := utils.ComparePassword(user.Password, strings.ToLower(user.Email)) // check if the password is valid
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password or email is incorrect",
		})
		return
	}

	token, err := utils.GenerateToken(strings.ToLower(user.Email)) // generate the JWT token

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error generating token",
		})
		return
	}

	database.Where("email = ?", strings.ToLower(user.Email)).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"userdata": gin.H{
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
