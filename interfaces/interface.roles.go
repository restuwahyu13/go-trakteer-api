package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IRolesRepository interface {
	CreateRepository(ctx context.Context, body *dtos.DTORoles) helpers.APIResponse
	GetAllRepository(ctx context.Context, query *dtos.DTORolePagination) helpers.APIResponse
	GetByIdRepository(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse
	DeleteByIdRepository(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse
	UpdatedByIdRepository(ctx context.Context, body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse
}

type IRolesService interface {
	CreateService(ctx context.Context, body *dtos.DTORoles) helpers.APIResponse
	GetAllService(ctx context.Context, query *dtos.DTORolePagination) helpers.APIResponse
	GetByIdService(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse
	DeleteByIdService(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse
	UpdatedByIdService(ctx context.Context, body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse
}

type IRolesController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetAllController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteByIdController(rw http.ResponseWriter, r *http.Request)
	UpdatedByIdController(rw http.ResponseWriter, r *http.Request)
}
