package list

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/selosele/nancy/models"
)

/* 파일 목록 조회 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 목록 조회 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// 파일 목록을 조회한다(PublicID가 없으면 전체 파일 목록을 조회).
	resp, err := h.Params.Cld.Admin.Asset(
		h.Params.Ctx,
		admin.AssetParams{PublicID: ""},
	)

	if err != nil {
		log.Fatalf("Failed to get list of files, %v\n", err)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Response)
}
