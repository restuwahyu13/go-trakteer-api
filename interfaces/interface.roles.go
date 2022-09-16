package interfaces

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IRolesRepository interface {
	CreateRepository(body *dtos.DTORoles) helpers.APIResponse
	GetAllRepository(query *dtos.DTORolePagination) helpers.APIResponse
	GetByIdRepository(params *dtos.DTORolesById) helpers.APIResponse
	DeleteByIdRepository(params *dtos.DTORolesById) helpers.APIResponse
	UpdatedByIdRepository(body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse
}

type IRolesService interface {
	CreateService(body *dtos.DTORoles) helpers.APIResponse
	GetAllService(query *dtos.DTORolePagination) helpers.APIResponse
	GetByIdService(params *dtos.DTORolesById) helpers.APIResponse
	DeleteByIdService(params *dtos.DTORolesById) helpers.APIResponse
	UpdatedByIdService(body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse
}

type IRolesController interface {
	CreateController(rw http.ResponseWriter, r *http.Request)
	GetAllController(rw http.ResponseWriter, r *http.Request)
	GetByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteByIdController(rw http.ResponseWriter, r *http.Request)
	UpdatedByIdController(rw http.ResponseWriter, r *http.Request)
}
