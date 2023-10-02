package folders

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/selosele/nancy/models"
)

/* 폴더 목록 조회 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 폴더 목록 조회 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// 폴더 목록을 조회한다.
	result, err := h.Params.Cld.Admin.SubFolders(
		h.Params.Ctx,
		admin.SubFoldersParams{
			Folder: r.FormValue("folder"), // folder 파라미터가 없으면 전체 폴더 목록을 조회
		},
	)

	if err != nil {
		log.Fatalf("Failed to get folders: %v\n", err)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Folders)
}
