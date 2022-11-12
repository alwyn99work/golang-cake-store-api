package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exeption"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/request"
	"golang-restful-api/model/response"
	"golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CakeService struct {
	CakeRepository repository.CakeRepositoryInterface
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCakeService(cakeRepository repository.CakeRepositoryInterface, DB *sql.DB, validate *validator.Validate) CakeServiceInterface {
	return &CakeService{
		CakeRepository: cakeRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CakeService) Create(ctx context.Context, request request.CakeCreateRequest) response.CakeResponseCreate {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	cake := domain.Cake{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
	}

	cake = service.CakeRepository.Save(ctx, tx, cake)

	return helper.ToCakeResponseCreate(cake)
}

func (service *CakeService) Update(ctx context.Context, request request.CakeUpdateRequest) response.CakeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	cake.Title = helper.RequestCheck(request.Title, cake.Title).(string)
	cake.Description = helper.RequestCheck(request.Description, cake.Description).(string)
	cake.Rating = helper.RequestCheck(request.Rating, cake.Rating).(float32)
	cake.Image = helper.RequestCheck(request.Image, cake.Image).(string)

	cake = service.CakeRepository.Update(ctx, tx, cake)

	return helper.ToCakeResponse(cake)
}

func (service *CakeService) Delete(ctx context.Context, cakeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	_, err = service.CakeRepository.FindById(ctx, tx, cakeId)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	service.CakeRepository.Delete(ctx, tx, cakeId)
}

func (service *CakeService) FindById(ctx context.Context, cakeId int) response.CakeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, cakeId)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	return helper.ToCakeResponse(cake)
}

func (service *CakeService) FindAll(ctx context.Context) []response.CakeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.CakeRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoriesResponse(categories)
}
