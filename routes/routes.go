package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/selosele/nancy/handlers/details"
	"github.com/selosele/nancy/handlers/list"
	"github.com/selosele/nancy/handlers/remove"
	"github.com/selosele/nancy/handlers/upload"
	"github.com/selosele/nancy/models"
)

/* 라우트 설정 */
func Setup(p models.HandlerParams) {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	// 파일 조회 Handler
	detailsHandler := details.Handler{Params: p}
	apiRouter.HandleFunc("/files/{publicId}", detailsHandler.HandleRequest).Methods(http.MethodGet)

	// 파일 목록 조회 Handler
	listHandler := list.Handler{Params: p}
	apiRouter.HandleFunc("/files", listHandler.HandleRequest).Methods(http.MethodGet)

	// 파일 삭제 Handler
	removeHandler := remove.Handler{Params: p}
	apiRouter.HandleFunc("/files/{publicId}", removeHandler.HandleRequest).Methods(http.MethodDelete)

	// 파일 업로드 Handler
	uploadHandler := upload.Handler{Params: p}
	apiRouter.HandleFunc("/files", uploadHandler.HandleRequest).Methods(http.MethodPost)

	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
