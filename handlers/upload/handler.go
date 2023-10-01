package upload

import (
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/selosele/nancy/models"
)

/* 파일 업로드 Handler */
func Handle(h models.Handler, w http.ResponseWriter) error {
	uploadResult, err := h.Cld.Upload.Upload(
		h.Ctx,
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
