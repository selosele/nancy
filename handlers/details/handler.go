package details

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

/* 파일 조회 Handler */
func Handle(w http.ResponseWriter, cld *cloudinary.Cloudinary, ctx context.Context) error {
	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "y2ofmteyrfkvg8jldraz.jpg"})

	if err != nil {
		log.Fatalf("Failed to get details of files, %v\n", err)
		return err
	}

	log.Println(resp)
	return nil
}
