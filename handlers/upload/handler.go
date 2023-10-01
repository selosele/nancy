package upload

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/selosele/nancy/models"
)

/* 파일 업로드 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 업로드 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	// HTTP 요청에서 파일 데이터를 받아온다.
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Cloudinary에 파일을 업로드한다.
	resp, err := h.Params.Cld.Upload.Upload(
		h.Params.Ctx,
		file,
		uploader.UploadParams{
			PublicID: r.FormValue("publicId"),
		},
	)

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	// 폴더명은 받아왔는데 파일명이 누락되었을 경우
	// 예) publicId: myfolder/
	if resp.SecureURL == "" {
		msg := "Failed to upload file: " + resp.SecureURL
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	log.Printf("Successfully uploaded file: %v", resp.SecureURL)

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Response)
}
