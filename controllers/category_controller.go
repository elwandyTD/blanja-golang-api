package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
