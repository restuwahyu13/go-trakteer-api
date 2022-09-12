package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type RolesService struct {
	repository *repositorys.RolesRepository
}

func NewRolesService(repository *repositorys.RolesRepository) *RolesService {
	return &RolesService{repository: repository}
}

func (ctx *RolesService) CreateService(body dtos.DTORoles) helpers.APIResponse {
	return ctx.repository.CreateRepository(body)
}

func (ctx *RolesService) GetAllService(query dtos.DTORolePagination) helpers.APIResponse {
	return ctx.repository.GetAllRepository(query)
}

func (ctx *RolesService) GetByIdService(params dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(params)
}

func (ctx *RolesService) DeleteByIdService(params dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(params)
}

func (ctx *RolesService) UpdatedByIdService(body dtos.DTORoles, params dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(body, params)
}
