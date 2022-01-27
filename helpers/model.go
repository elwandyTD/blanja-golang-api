package helpers

import (
	"github.com/elwandyTD/blanja-golang-api/models/domain"
	"github.com/elwandyTD/blanja-golang-api/models/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Name:       category.Name,
		Image:      category.Image,
		CreateTime: category.CreateTime,
		UpdateTime: category.UpdateTime,
		Active:     category.Active,
	}
}
