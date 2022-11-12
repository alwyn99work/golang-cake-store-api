package router

import (
	"golang-restful-api/controller"
	"golang-restful-api/exeption"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cakeController controller.CakeControllerInterface) *httprouter.Router {

	router := httprouter.New()

	api := "/api"
	router.GET(api+"/cakes", cakeController.FindAll)
	router.GET(api+"/cakes/:id", cakeController.FindById)
	router.POST(api+"/cakes", cakeController.Create)
	router.PUT(api+"/cakes/:id", cakeController.Update)
	router.DELETE(api+"/cakes/:id", cakeController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	return router
}
