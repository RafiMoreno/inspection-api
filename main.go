package main

import (
	"inspection-api/initializers"
	"inspection-api/services"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.SetupDatabase()
	initializers.SetupCloudinary()
	
	if initializers.DB == nil {
        log.Fatal("Database connection is not initialized")
    }

	initializers.MigrateDB(initializers.DB)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FE_URL")}, 
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	endpoint := r.Group("/api")
	{
		endpoint.POST("/upload", services.UploadImages)
	}

	r.Run(":8080") 
}
