package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckParse(c *gin.Context, model interface{}) interface{} {
	parsing := c.BindJSON(&model)
	if parsing != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error parsing JSON",
		})
		return nil
	}

	return model
}
