package upload

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

/* 파일 업로드 Handler */
func Handle(w http.ResponseWriter, cld *cloudinary.Cloudinary, ctx context.Context) error {
	uploadResult, err := cld.Upload.Upload(
		ctx,
		"https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png",
		uploader.UploadParams{PublicID: "logo"},
	)

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return err
	}

	log.Printf("Succesed to upload file: %v", uploadResult.SecureURL)
	return nil
}
