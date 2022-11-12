package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/response"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("X-API-Key")
	if header == "RAHASIA" {
		// ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		apiResponse := response.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, apiResponse)
	}
}
