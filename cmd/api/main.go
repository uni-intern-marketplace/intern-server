package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uni-intern-marketplace/intern-server/internal/database"
)

func main() {
	// Базаға қосылу
	dsn := os.Getenv("DATABASE_URL")
	database.ConnectDatabase(dsn)

	r := gin.Default()

	// CORS баптау (Next.js-пен байланыс үшін маңызды)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.js порты
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Тесттік роут
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP", "message": "Go server is running"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
