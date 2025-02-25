package services

import (
	"context"
	"inspection-api/initializers"
	"inspection-api/models"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadToCloudinary(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "Error encoutered", err
	}
	defer file.Close()

	uploadResult, err := initializers.Cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "Error encoutered", err
	}

	return uploadResult.SecureURL, nil
}

func isImageFile(fileHeader *multipart.FileHeader) bool {
	validExt := map[string]bool {
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
		".avif": true,
		".svg":  true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	return validExt[ext]
}

func UploadImages(c *gin.Context) {
	form, _ := c.MultipartForm()
	images := form.File["images"]
	labels := form.Value["labels"]

	//Handle missing values
	if len(images) != len(labels) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing values"})
		return
	}

	//Handle non-image files
	for _, image := range images {
		if !isImageFile(image) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Uncompatible format"})
			return
		}
	}
	
	var listOfImages []models.ImageField
	for i, image := range images {
		imageURL, err := UploadToCloudinary(image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Upload Image"})
			return
		}

		imageField := models.ImageField{Label: labels[i], ImageUrl: imageURL}
		initializers.DB.Create(&imageField)

		listOfImages = append(listOfImages, imageField)
	}

	c.JSON(http.StatusOK, gin.H{"Uploaded Images": listOfImages})
	
}