package main

import (
	"inspection-api/initializers"
	"inspection-api/services"
	"log"

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

	endpoint := r.Group("/api")
	{
		endpoint.POST("/upload", services.UploadImages)
	}

	r.Run(":8080") 
}
