package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IDonationTypeRepository interface {
	CreateRepository(ctx context.Context, body *dtos.DTODonationType) helpers.APIResponse
	GetAllRepository(ctx context.Context, query *dtos.DTODonationTypePagination) helpers.APIResponse
	GetByIdRepository(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse
	DeleteByIdRepository(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse
	UpdateByIdRepository(ctx context.Context, body *dtos.DTODonationType, params *dtos.DTODonationTypeId) helpers.APIResponse
}

type IDonationTypeService interface {
	CreateService(ctx context.Context, body *dtos.DTODonationType) helpers.APIResponse
	GetAllService(ctx context.Context, query *dtos.DTODonationTypePagination) helpers.APIResponse
	GetByIdService(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse
	DeleteByIdService(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse
	UpdateByIdService(ctx context.Context, body *dtos.DTODonationType, params *dtos.DTODonationTypeId) helpers.APIResponse
}

type IDonationTypeController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetAllController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteByIdController(rw http.ResponseWriter, r *http.Request)
	UpdatedByIdController(rw http.ResponseWriter, r *http.Request)
}
