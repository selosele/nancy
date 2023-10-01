package handlers

import (
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gorilla/mux"
	"github.com/selosele/nancy/handlers/details"
	"github.com/selosele/nancy/handlers/list"
	"github.com/selosele/nancy/handlers/upload"
)

/* 라우트 설정 */
func Setup(cld *cloudinary.Cloudinary, ctx context.Context) {
	r := mux.NewRouter()

	// 파일 조회 Handler
	r.HandleFunc("/files/{public_id}", func(w http.ResponseWriter, r *http.Request) {
		details.Handle(w, cld, ctx)
	}).Methods(http.MethodGet)

	// 파일 목록 조회 Handler
	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		list.Handle(w, cld, ctx)
	}).Methods(http.MethodGet)

	// 파일 업로드 Handler
	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		upload.Handle(w, cld, ctx)
	}).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
