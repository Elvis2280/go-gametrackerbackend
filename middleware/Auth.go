package middleware

import (
	"gametracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization") // get the token from the header

		if token == "" { // if there is no token, return error
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No token provided",
			})

			c.Abort()
			return
		}

		extractedToken := strings.Split(token, "Bearer ") // split the token

		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})

			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(token) // verify the token

		if err != nil { // if there is an error, return error
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})

			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()

	}
}
