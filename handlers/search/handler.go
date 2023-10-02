package search

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/selosele/nancy/models"
	"github.com/selosele/nancy/utils"
)

/* 파일 검색 Handler 구조체 */
type Handler struct {
	Params models.HandlerParams
}

/* 파일 검색 HTTP 요청 처리 */
func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	resourceType := utils.DefaultString(r.FormValue("resourceType"), "image")
	publicId := r.FormValue("publicId")
	folder := r.FormValue("folder")
	maxResults, _ := strconv.Atoi(utils.DefaultString(r.FormValue("maxResults"), "30"))
	expression := "resource_type:" + resourceType

	// publicId로 검색한다.
	if utils.IsNotBlank(publicId) {
		expression += " AND public_id:" + publicId
	}

	// folder로 검색한다.
	if utils.IsNotBlank(folder) {
		expression += " AND folder:" + folder + "/*"
	}

	searchQuery := search.Query{
		Expression: expression,
		SortBy:     []search.SortByField{{"public_id": search.Ascending}},
		MaxResults: maxResults,
	}

	// 파일을 검색한다.
	result, err := h.Params.Cld.Admin.Search(
		h.Params.Ctx,
		searchQuery,
	)

	if err != nil {
		log.Fatalf("Failed to search files: %v\n", err)
	}

	// JSON을 응답으로 보낸다.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Response)
}
