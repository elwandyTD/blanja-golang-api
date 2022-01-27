package controllers

import (
	"context"
	"net/http"

	"github.com/elwandyTD/blanja-golang-api/helpers"
	"github.com/elwandyTD/blanja-golang-api/models/web"
	"github.com/elwandyTD/blanja-golang-api/services"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	// id, err := strconv.Atoi(categoryId)
	// helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(context.Background(), categoryId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(writer, webResponse)
}
