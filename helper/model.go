package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/response"
)

func ToCakeResponse(cake domain.Cake) response.CakeResponse {
	return response.CakeResponse{
		Id:          cake.Id,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   cake.CreatedAt,
		UpdatedAt:   cake.UpdatedAt,
	}
}

func ToCakeResponseCreate(cake domain.Cake) response.CakeResponseCreate {
	return response.CakeResponseCreate{
		Id:          cake.Id,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}
}

func ToCategoriesResponse(categories []domain.Cake) []response.CakeResponse {
	var cakeResponses []response.CakeResponse
	for _, cake := range categories {
		cakeResponses = append(cakeResponses, ToCakeResponse(cake))
	}

	return cakeResponses
}
