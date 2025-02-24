package main

import (
	"inspection-api/initializers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.SetupDatabase()
	
	if initializers.DB == nil {
        log.Fatal("Database connection is not initialized")
    }

	initializers.MigrateDB(initializers.DB)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Inspection API!",
		})
	})

	r.Run(":8080") 
}
