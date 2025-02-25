package initializers

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cld *cloudinary.Cloudinary

func SetupCloudinary() {
	var err error
	Cld, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatal("Cloudinary setup failed", err)
	}
}