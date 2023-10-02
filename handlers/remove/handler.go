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
		log.Fatalf("Failed to remove file: %v\n", err)
	}

	var resp any

	// 오류 발생 시, 오류 메시지를 출력한다.
	if utils.IsNotBlank(result.Error.Message) {
		log.Printf("Failed to remove file: %v\n", result.Error.Message)
		resp = result.Error
	} else {

		if result.Result != "ok" { // 파일 삭제 실패 시
			log.Printf("Failed to remove file: %v\n", result.Result)
			resp = result.Response
		} else { // 파일 삭제 성공 시
			resp = result.Response
			log.Printf("Successfully removed file: %v\n", result.Result)
		}
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
