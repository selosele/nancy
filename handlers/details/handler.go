package details

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

/* 파일 조회 Handler */
func Handle(cld *cloudinary.Cloudinary) error {
	ctx := context.Background()

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "my_image"})

	if err != nil {
		log.Fatalf("Failed to get details of files, %v\n", err)
		return err
	}

	log.Println(resp.SecureURL)
	return nil
}
