package main

import (
	"net/http"

	"github.com/elwandyTD/blanja-golang-api/app"
	"github.com/elwandyTD/blanja-golang-api/controllers"
	"github.com/elwandyTD/blanja-golang-api/helpers"
	"github.com/elwandyTD/blanja-golang-api/repositories"
	"github.com/elwandyTD/blanja-golang-api/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := "localhost:3000"

	db := app.NewDB()
	validate := validator.New()

	helpers.ViperInitConfigLogging()

	logrus.WithField("addr", addr).Info("Host name")

	categoryRepositories := repositories.NewCategoryRepository()
	categoryServices := services.NewCategoryService(categoryRepositories, db, validate)
	categoryController := controllers.NewCategoryController(categoryServices)

	router := httprouter.New()

	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
