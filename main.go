package main

import (
	"log"
	"room-reserve/db"
	"room-reserve/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/cors"
)

func main() {
	db.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("FAIILED TO START SERVER ⛔")
	}
}
