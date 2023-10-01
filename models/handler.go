package models

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

/* Handler 구조체 */
type Handler struct {
	Cld *cloudinary.Cloudinary
	Ctx context.Context
}
