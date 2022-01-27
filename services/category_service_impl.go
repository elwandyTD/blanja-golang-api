package services

import (
	"context"
	"database/sql"

	"github.com/elwandyTD/blanja-golang-api/helpers"
	"github.com/elwandyTD/blanja-golang-api/models/web"

	// "github.com/elwandyTD/blanja-golang-api/models/domain"
	"github.com/elwandyTD/blanja-golang-api/repositories"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
// 	tx, err := service.DB.Begin()
// 	helpers.PanicIfError(err)
// 	defer helpers.CommitOrRollback(tx)

// 	// category := domain.Category{
// 	// 	Name:  request.Name,
// 	// 	Image: "Test",
// 	// }

// 	// return helpers.ToCategoryResponse(domain.Category())
// }

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(err.Error())
	}

	return helpers.ToCategoryResponse(category)
}
