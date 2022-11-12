package main

import (
	"golang-restful-api/controller"
	"golang-restful-api/database"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/router"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {
	mysqlDB := database.NewMysqlDB()
	database.Migrate(mysqlDB)
	validate := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, mysqlDB, validate)
	cakeController := controller.NewCakeController(cakeService)

	router := router.NewRouter(cakeController)

	handler := cors.AllowAll().Handler(middleware.NewAuthMiddleware(router))
	server := http.Server{
		Addr:    "localhost:8001",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
