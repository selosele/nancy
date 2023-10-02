package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/selosele/nancy/handlers/remove"
	"github.com/selosele/nancy/handlers/upload"
	"github.com/selosele/nancy/handlers/view"
	"github.com/selosele/nancy/models"
)

/* 라우트 설정 */
func Setup(p models.HandlerParams) {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	// 파일 조회 Handler
	viewHandler := view.Handler{Params: p}
	apiRouter.HandleFunc("/files", viewHandler.HandleRequest).Methods(http.MethodGet)

	// 파일 업로드 Handler
	uploadHandler := upload.Handler{Params: p}
	apiRouter.HandleFunc("/files", uploadHandler.HandleRequest).Methods(http.MethodPost)

	// 파일 삭제 Handler
	removeHandler := remove.Handler{Params: p}
	apiRouter.HandleFunc("/files", removeHandler.HandleRequest).Methods(http.MethodDelete)

	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
