package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type CategoriesService = interfaces.ICategoriesService
type categoriesService struct {
	repository repositorys.CategoriesRepository
}

func NewCategoriesService(repository repositorys.CategoriesRepository) *categoriesService {
	return &categoriesService{repository: repository}
}

func (ctx *categoriesService) CreateService(body *dtos.DTOCategories) helpers.APIResponse {
	return ctx.repository.CreateRepository(body)
}

func (ctx *categoriesService) GetAllService(query *dtos.DTOCategoriesPagination) helpers.APIResponse {
	return ctx.repository.GetAllRepository(query)
}

func (ctx *categoriesService) GetByIdService(params *dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(params)
}

func (ctx *categoriesService) DeleteByIdService(params *dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(params)
}

func (ctx *categoriesService) UpdatedByIdService(body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(body, params)
}
