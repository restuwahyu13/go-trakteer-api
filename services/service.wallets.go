package services

import (
	"context"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type WalletsService = interfaces.IWalletsService
type walletsService struct {
	repository repositorys.WalletsRepository
}

func NewWalletsService(repository repositorys.WalletsRepository) *walletsService {
	return &walletsService{repository: repository}
}

func (s *walletsService) CreateService(ctx context.Context, body *dtos.DTOWalletsCreate) helpers.APIResponse {
	res := s.repository.CreateRepository(ctx, body)
	return res
}

func (s *walletsService) GetByIdService(ctx context.Context, params *dtos.DTOWalletsById) helpers.APIResponse {
	res := s.repository.GetByIdRepository(ctx, params)
	return res
}

func (s *walletsService) UpdateByIdService(ctx context.Context, body *dtos.DTOWalletsUpdate, params *dtos.DTOWalletsById) helpers.APIResponse {
	res := s.repository.UpdateByIdRepository(ctx, body, params)
	return res
}
