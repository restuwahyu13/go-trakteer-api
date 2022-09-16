package interfaces

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type ICategoriesRepository interface {
	CreateRepository(body *dtos.DTOCategories) helpers.APIResponse
	GetAllRepository(query *dtos.DTOCategoriesPagination) helpers.APIResponse
	GetByIdRepository(params *dtos.DTOCategoriesId) helpers.APIResponse
	DeleteByIdRepository(params *dtos.DTOCategoriesId) helpers.APIResponse
	UpdatedByIdRepository(body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse
}

type ICategoriesService interface {
	CreateService(body *dtos.DTOCategories) helpers.APIResponse
	GetAllService(query *dtos.DTOCategoriesPagination) helpers.APIResponse
	GetByIdService(params *dtos.DTOCategoriesId) helpers.APIResponse
	DeleteByIdService(params *dtos.DTOCategoriesId) helpers.APIResponse
	UpdatedByIdService(body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse
}

type ICategoriesController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetAllController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteByIdController(rw http.ResponseWriter, r *http.Request)
	UpdatedByIdController(rw http.ResponseWriter, r *http.Request)
}
