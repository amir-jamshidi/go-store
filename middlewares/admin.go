package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists := c.Get("user_role")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not authenticated",
			})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Admin role required",
			})
		}

		c.Next()
	}

}
