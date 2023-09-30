package handlers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/selosele/nancy/handlers/details"
	"github.com/selosele/nancy/handlers/upload"
)

/* 라우트 설정 */
func Init(cld *cloudinary.Cloudinary) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 이미지 조회 Handler
		if r.RequestURI == "/details" && r.Method == http.MethodGet {
			details.Handle(cld)
			return
		}

		// 이미지 업로드 Handler
		if r.RequestURI == "/upload" && r.Method == http.MethodPost {
			upload.Handle(cld)
			return
		}

		// 요청이 지원되지 않는 경우 405 Method Not Allowed 응답을 보냄
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}
