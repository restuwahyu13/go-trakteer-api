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

func (ctx *CategoriesService) CreateService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.CreateRepository(payload)
}

func (ctx *CategoriesService) GetAllService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetAllRepository(payload)
}

func (ctx *CategoriesService) GetByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(payload)
}

func (ctx *CategoriesService) DeleteByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(payload)
}

func (ctx *CategoriesService) UpdatedByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(payload)
}
