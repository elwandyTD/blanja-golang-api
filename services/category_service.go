package services

import (
	"context"

	"github.com/elwandyTD/blanja-golang-api/models/web"
)

type CategoryService interface {
	// Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindById(ctx context.Context, categoryId string) web.CategoryResponse
}
