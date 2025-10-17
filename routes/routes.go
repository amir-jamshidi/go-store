package routes

import (
	"room-reserve/handlers"
	"room-reserve/repositories"
	"room-reserve/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository()
	otpRepo := repositories.NewOTPRepository()

	// Initialize services
	authService := services.NewAuthService(userRepo, otpRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.POST("/send-otp", authHandler.SendOTP)
			users.POST("/verify-otp", authHandler.VerifyOTP)
		}
	}

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
