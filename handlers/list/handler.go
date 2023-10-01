package list

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/selosele/nancy/models"
)

/* 파일 목록 조회 Handler */
func Handle(h models.Handler, w http.ResponseWriter) error {
	resp, err := h.Cld.Admin.Asset(h.Ctx, admin.AssetParams{PublicID: ""}) // public_id가 없으면 전체 파일 목록을 조회

	if err != nil {
		log.Fatalf("Failed to get list of files, %v\n", err)
		return err
	}

	// 응답 결과를 JSON 형식의 바이트 슬라이스로 변환
	json, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal JSON, %v\n", err)
		return err
	}

	w.Write(json)
	return nil
}
