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

func (ctx *RolesService) CreateService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.CreateRepository(payload)
}

func (ctx *RolesService) GetAllService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetAllRepository(payload)
}

func (ctx *RolesService) GetByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetByIdRepository(payload)
}

func (ctx *RolesService) DeleteByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.DeleteByIdRepository(payload)
}

func (ctx *RolesService) UpdatedByIdService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.UpdatedByIdRepository(payload)
}
