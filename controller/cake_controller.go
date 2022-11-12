package controller

import (
	"encoding/json"
	"golang-restful-api/helper"
	"golang-restful-api/model/request"
	"golang-restful-api/model/response"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CakeController struct {
	CakeService service.CakeServiceInterface
}

func NewCakeController(cakeService service.CakeServiceInterface) CakeControllerInterface {
	return &CakeController{
		CakeService: cakeService,
	}
}

func (controller *CakeController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.Log("POST /api/cakes")

	// format request
	decoder := json.NewDecoder(r.Body)
	cakeCreateRequest := request.CakeCreateRequest{}
	err := decoder.Decode(&cakeCreateRequest)
	helper.PanicIfError(err)

	// format data
	cakeResponse := controller.CakeService.Create(r.Context(), cakeCreateRequest)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",
		Data:   cakeResponse,
	}

	// format response
	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *CakeController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.Log("PUT /api/cakes/{id}")

	// format request
	cakeUpdateRequest := request.CakeUpdateRequest{}
	helper.ReadFromRequestBody(r, &cakeUpdateRequest)

	// convert id from type string to int first
	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.PanicIfError(err)
	cakeUpdateRequest.Id = id

	// format data
	cakeResponse := controller.CakeService.Update(r.Context(), cakeUpdateRequest)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",
		Data:   cakeResponse,
	}

	// format response
	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *CakeController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.Log("DELETE /api/cakes/{id}")

	// convert id from type string to int first
	cakeId := params.ByName("id")
	id := helper.TransformId(cakeId)

	// format data
	controller.CakeService.Delete(r.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",
	}

	// format response
	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *CakeController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.Log("GET /api/cakes/{id}")

	// convert id from type string to int first
	cakeId := params.ByName("id")
	id := helper.TransformId(cakeId)

	// format data
	cakeResponse := controller.CakeService.FindById(r.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",
		Data:   cakeResponse,
	}

	// format response
	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *CakeController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helper.Log("GET /api/cakes")

	// format data
	cakesResponse := controller.CakeService.FindAll(r.Context())
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",
		Data:   cakesResponse,
	}

	// format response
	helper.WriteToResponseBody(w, apiResponse)
}
