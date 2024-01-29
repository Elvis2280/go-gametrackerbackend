package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckParse(c *gin.Context, model interface{}) interface{} {

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return nil
	}

	return model
}
