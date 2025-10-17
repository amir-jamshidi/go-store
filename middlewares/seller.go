package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SellerRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not authenticated",
			})
			c.Abort()
			return
		}

		if role != "admin" && role != "seller" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Seller role required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
