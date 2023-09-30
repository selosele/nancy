package list

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

/* 파일 목록 조회 Handler */
func Handle(w http.ResponseWriter, cld *cloudinary.Cloudinary, ctx context.Context) error {
	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: ""}) // public_id가 없으면 전체 파일 목록을 조회

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
