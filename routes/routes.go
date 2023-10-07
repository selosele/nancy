package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/selosele/nancy/handlers/folders"
	"github.com/selosele/nancy/handlers/remove"
	"github.com/selosele/nancy/handlers/search"
	"github.com/selosele/nancy/handlers/upload"
	"github.com/selosele/nancy/models"
)

/* 라우트 설정 */
func Setup(p models.HandlerParams) {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	// 파일 검색 Handler
	searchHandler := search.Handler{Params: p}
	apiRouter.HandleFunc("/files", searchHandler.HandleRequest).Methods("GET")

	// 파일 업로드 Handler
	uploadHandler := upload.Handler{Params: p}
	apiRouter.HandleFunc("/files", uploadHandler.HandleRequest).Methods("POST")

	// 파일 삭제 Handler
	removeHandler := remove.Handler{Params: p}
	apiRouter.HandleFunc("/files", removeHandler.HandleRequest).Methods("DELETE")

	// 폴더 목록 조회 Handler
	foldersHandler := folders.Handler{Params: p}
	apiRouter.HandleFunc("/folders", foldersHandler.HandleRequest).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
