package service

import (
	"context"
	"golang-restful-api/model/request"
	"golang-restful-api/model/response"
)

type CakeServiceInterface interface {
	Create(ctx context.Context, request request.CakeCreateRequest) response.CakeResponseCreate
	Update(ctx context.Context, request request.CakeUpdateRequest) response.CakeResponse
	Delete(ctx context.Context, cakeId int)
	FindById(ctx context.Context, cakeId int) response.CakeResponse
	FindAll(ctx context.Context) []response.CakeResponse
}
