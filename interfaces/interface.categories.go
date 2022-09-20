package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type ICategoriesRepository interface {
	CreateRepository(ctx context.Context, body *dtos.DTOCategories) helpers.APIResponse
	GetAllRepository(ctx context.Context, query *dtos.DTOCategoriesPagination) helpers.APIResponse
	GetByIdRepository(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse
	DeleteByIdRepository(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse
	UpdatedByIdRepository(ctx context.Context, body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse
}

type ICategoriesService interface {
	CreateService(ctx context.Context, body *dtos.DTOCategories) helpers.APIResponse
	GetAllService(ctx context.Context, query *dtos.DTOCategoriesPagination) helpers.APIResponse
	GetByIdService(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse
	DeleteByIdService(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse
	UpdatedByIdService(ctx context.Context, body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse
}

type ICategoriesController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetAllController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteByIdController(rw http.ResponseWriter, r *http.Request)
	UpdatedByIdController(rw http.ResponseWriter, r *http.Request)
}
