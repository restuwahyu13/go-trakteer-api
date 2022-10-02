package services

import (
	"context"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type DonationTypeService = interfaces.IDonationTypeService
type donationTypeService struct {
	repository repositorys.DonationTypeRepository
}

func NewDonationTypeService(repository repositorys.DonationTypeRepository) *donationTypeService {
	return &donationTypeService{repository: repository}
}

func (s *donationTypeService) CreateService(ctx context.Context, body *dtos.DTODonationType) helpers.APIResponse {
	return s.repository.CreateRepository(ctx, body)
}

func (s *donationTypeService) GetAllService(ctx context.Context, query *dtos.DTODonationTypePagination) helpers.APIResponse {
	return s.repository.GetAllRepository(ctx, query)
}

func (s *donationTypeService) GetByIdService(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse {
	return s.repository.GetByIdRepository(ctx, params)
}

func (s *donationTypeService) DeleteByIdService(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse {
	return s.repository.DeleteByIdRepository(ctx, params)
}

func (s *donationTypeService) UpdateByIdService(ctx context.Context, body *dtos.DTODonationType, params *dtos.DTODonationTypeId) helpers.APIResponse {
	return s.repository.UpdateByIdRepository(ctx, body, params)
}
