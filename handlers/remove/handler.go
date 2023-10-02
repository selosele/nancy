package remove

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/selosele/nancy/models"
	"github.com/selosele/nancy/utils"
)

/* 파일 삭제 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 삭제 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// Cloudinary의 파일을 삭제한다.
	result, err := h.Params.Cld.Upload.Destroy(
		h.Params.Ctx,
		uploader.DestroyParams{
			PublicID:     r.FormValue("publicId"),
			ResourceType: utils.DefaultString(r.FormValue("resourceType"), "image"),
		},
	)

	if err != nil {
		log.Fatalf("Failed to remove file, %v\n", err)
	}

	// 삭제하고자 하는 파일이 없는 경우
	if result.Result == "not found" {
		log.Printf("Failed to remove file, %v\n", result.Result)
	} else {
		log.Printf("Successfully removed file: %v\n", result.Response)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Response)
}
