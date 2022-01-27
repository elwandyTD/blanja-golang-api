package main

import (
	"log"
	"net/http"

	"github.com/elwandyTD/blanja-golang-api/app"
	"github.com/elwandyTD/blanja-golang-api/controllers"
	"github.com/elwandyTD/blanja-golang-api/helpers"
	"github.com/elwandyTD/blanja-golang-api/repositories"
	"github.com/elwandyTD/blanja-golang-api/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepositories := repositories.NewCategoryRepository()
	categoryServices := services.NewCategoryService(categoryRepositories, db, validate)
	categoryController := controllers.NewCategoryController(categoryServices)

	router := httprouter.New()

	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)

	log.Print("Ready")
	log.Printf("Ready")
	log.Println("Ready")
}
