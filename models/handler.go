package models

import (
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
)

/* HTTP Handler 인터페이스 */
type Handler interface {

	/* HTTP 요청을 처리한다. */
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

/* Handler 매개변수 구조체 */
type HandlerParams struct {
	Cld *cloudinary.Cloudinary
	Ctx context.Context
}
