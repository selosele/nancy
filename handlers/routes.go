package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/selosele/nancy/handlers/details"
	"github.com/selosele/nancy/handlers/list"
	"github.com/selosele/nancy/handlers/upload"
	"github.com/selosele/nancy/models"
)

/* 라우트 설정 */
func Setup(h models.Handler) {
	r := mux.NewRouter()

	// 파일 조회 Handler
	r.HandleFunc("/files/{public_id}", func(w http.ResponseWriter, r *http.Request) {
		details.Handle(h, w)
	}).Methods(http.MethodGet)

	// 파일 목록 조회 Handler
	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		list.Handle(h, w)
	}).Methods(http.MethodGet)

	// 파일 업로드 Handler
	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		upload.Handle(h, w)
	}).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
