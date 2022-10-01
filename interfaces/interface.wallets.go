package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IWalletsRepository interface {
	CreateRepository(ctx context.Context, body *dtos.DTOWalletsCreate) helpers.APIResponse
	GetByIdRepository(ctx context.Context, params *dtos.DTOWalletsById) helpers.APIResponse
	UpdateByIdRepository(ctx context.Context, body *dtos.DTOWalletsUpdate, params *dtos.DTOWalletsById) helpers.APIResponse
}

type IWalletsService interface {
	CreateService(ctx context.Context, body *dtos.DTOWalletsCreate) helpers.APIResponse
	GetByIdService(ctx context.Context, params *dtos.DTOWalletsById) helpers.APIResponse
	UpdateByIdService(ctx context.Context, body *dtos.DTOWalletsUpdate, params *dtos.DTOWalletsById) helpers.APIResponse
}

type IWalletsController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	UpdateByIdController(rw http.ResponseWriter, r *http.Request)
}
