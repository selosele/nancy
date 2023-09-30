package upload

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

/* 파일 업로드 Handler */
func Handle(cld *cloudinary.Cloudinary) error {
	ctx := context.Background()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		"https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png",
		uploader.UploadParams{PublicID: "logo"})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return err
	}

	log.Println(uploadResult.SecureURL)
	return nil
}
