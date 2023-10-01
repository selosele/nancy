package details

import (
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/selosele/nancy/models"
)

/* 파일 조회 Handler */
func Handle(h models.Handler, w http.ResponseWriter) error {
	resp, err := h.Cld.Admin.Asset(h.Ctx, admin.AssetParams{PublicID: "y2ofmteyrfkvg8jldraz.jpg"})

	if err != nil {
		log.Fatalf("Failed to get details of files, %v\n", err)
		return err
	}

	log.Println(resp)
	return nil
}
