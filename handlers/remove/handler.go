package remove

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
	"github.com/selosele/nancy/models"
)

/* 파일 삭제 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 삭제 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// resourceType 파라미터가 비어 있는지 확인한다.
	if r.FormValue("resourceType") == "" {
		msg := "resourceType must be required"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// URI에서 {publicId} 값을 추출한다.
	vars := mux.Vars(r)
	publicId := r.FormValue("dir") + vars["publicId"]

	// Cloudinary의 파일을 삭제한다.
	resp, err := h.Params.Cld.Upload.Destroy(
		h.Params.Ctx,
		uploader.DestroyParams{
			PublicID:     publicId,
			ResourceType: r.FormValue("resourceType"),
		},
	)

	if err != nil {
		log.Fatalf("Failed to remove file, %v\n", err)
	}

	// 삭제하고자 하는 파일이 없는 경우
	if resp.Result == "not found" {
		log.Printf("Failed to remove file, %v\n", resp.Result)
	} else {
		log.Printf("Successfully removed file: %v\n", resp.Response)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Response)
}
