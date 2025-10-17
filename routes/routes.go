package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.POST("/send-otp", sendOTP)
			users.POST("/verify-otp", verifyOTP)
		}
	}

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
