package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type RolesService = interfaces.IRolesService
type rolesService struct {
	repository repositorys.RolesRepository
}

func NewRolesService(repository repositorys.RolesRepository) *rolesService {
	return &rolesService{repository: repository}
}

func (ctx *rolesService) CreateService(body *dtos.DTORoles) helpers.APIResponse {
	return ctx.repository.CreateRepository(body)
}

func (ctx *rolesService) GetAllService(query *dtos.DTORolePagination) helpers.APIResponse {
	return ctx.repository.GetAllRepository(query)
}

func (ctx *rolesService) GetByIdService(params *dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(params)
}

func (ctx *rolesService) DeleteByIdService(params *dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(params)
}

func (ctx *rolesService) UpdatedByIdService(body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(body, params)
}
