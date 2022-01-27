package web

import (
	"mime/multipart"
)

type CategoryCreateRequest struct {
	Name  string         `json:"name"`
	Image multipart.File `json:"image"`
}
