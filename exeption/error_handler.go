package exeption

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if validationError(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exeption, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		apiResponse := response.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exeption.Error(),
		}

		helper.WriteToResponseBody(w, apiResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	_, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		apiResponse := response.ApiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err,
		}

		helper.WriteToResponseBody(w, apiResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	apiResponse := response.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
