package services

import (
	"context"

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

func (s *rolesService) CreateService(ctx context.Context, body *dtos.DTORoles) helpers.APIResponse {
	return s.repository.CreateRepository(ctx, body)
}

func (s *rolesService) GetAllService(ctx context.Context, query *dtos.DTORolePagination) helpers.APIResponse {
	return s.repository.GetAllRepository(ctx, query)
}

func (s *rolesService) GetByIdService(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse {
	return s.repository.GetByIdRepository(ctx, params)
}

func (s *rolesService) DeleteByIdService(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse {
	return s.repository.DeleteByIdRepository(ctx, params)
}

func (s *rolesService) UpdatedByIdService(ctx context.Context, body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse {
	return s.repository.UpdatedByIdRepository(ctx, body, params)
}
