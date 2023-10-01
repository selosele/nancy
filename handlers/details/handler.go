package details

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/gorilla/mux"
	"github.com/selosele/nancy/models"
)

/* 파일 조회 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 조회 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	assetType := r.FormValue("assetType")

	// assetType 파라미터가 비어 있는지 확인한다.
	if assetType == "" {
		msg := "assetType must be required"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// URI에서 {publicId} 값을 추출한다.
	vars := mux.Vars(r)
	publicId := r.FormValue("dir") + vars["publicId"]

	// publicId로 파일을 조회한다.
	resp, err := h.Params.Cld.Admin.Asset(
		h.Params.Ctx,
		admin.AssetParams{
			PublicID:  publicId,
			AssetType: api.AssetType(assetType),
		},
	)

	if err != nil {
		log.Fatalf("Failed to get details of files, %v\n", err)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Response)
}
