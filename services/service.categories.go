package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type CategoriesService struct {
	repository *repositorys.CategoriesRepository
}

func NewCategoriesService(repository *repositorys.CategoriesRepository) *CategoriesService {
	return &CategoriesService{repository: repository}
}

func (ctx *CategoriesService) CreateService(body dtos.DTOCategories) helpers.APIResponse {
	return ctx.repository.CreateRepository(body)
}

func (ctx *CategoriesService) GetAllService(query dtos.DTOCategoriesPagination) helpers.APIResponse {
	return ctx.repository.GetAllRepository(query)
}

func (ctx *CategoriesService) GetByIdService(params dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(params)
}

func (ctx *CategoriesService) DeleteByIdService(params dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(params)
}

func (ctx *CategoriesService) UpdatedByIdService(body dtos.DTOCategories, params dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(body, params)
}
