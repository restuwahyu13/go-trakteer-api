package interfaces

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IUsersRepository interface {
	LoginRepository(body *dtos.DTOUsersLogin) helpers.APIResponse
	ForgotPasswordRepository(body *dtos.DTOUsersForgotPassword) helpers.APIResponse
	ResetPasswordRepository(body *dtos.DTOUsersResetPassword) helpers.APIResponse
	ChangePasswordRepository(body *dtos.DTOUsersChangePassword) helpers.APIResponse
	GetProfileByIdRepository(params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	UpdateProfileByIdRepository(body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	CreateUsersRepository(body *dtos.DTOUsersCreate) helpers.APIResponse
	GetAllUsersRepository(query *dtos.DTOUsersPagination) helpers.APIResponse
	GetUsersByIdRepository(params *dtos.DTOUsersById) helpers.APIResponse
	DeleteUsersByIdRepository(params *dtos.DTOUsersById) helpers.APIResponse
	UpdateUsersByIdRepository(body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse
}

type IUsersService interface {
	LoginService(body *dtos.DTOUsersLogin) helpers.APIResponse
	ForgotPasswordService(body *dtos.DTOUsersForgotPassword) helpers.APIResponse
	ResetPasswordService(body *dtos.DTOUsersResetPassword) helpers.APIResponse
	ChangePasswordService(body *dtos.DTOUsersChangePassword) helpers.APIResponse
	GetProfileByIdService(params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	UpdateProfileByIdService(body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	CreateUsersService(body *dtos.DTOUsersCreate) helpers.APIResponse
	GetAllUsersService(query *dtos.DTOUsersPagination) helpers.APIResponse
	GetUsersByIdService(params *dtos.DTOUsersById) helpers.APIResponse
	DeleteUsersByIdService(params *dtos.DTOUsersById) helpers.APIResponse
	UpdateUsersByIdService(body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse
}

type IUsersController interface {
	LoginController(rw http.ResponseWriter, r *http.Request)
	ForgotPasswordController(rw http.ResponseWriter, r *http.Request)
	ResetPasswordController(rw http.ResponseWriter, r *http.Request)
	ChangePasswordController(rw http.ResponseWriter, r *http.Request)
	GetProfileByIdController(rw http.ResponseWriter, r *http.Request)
	UpdateProfileByIdController(rw http.ResponseWriter, r *http.Request)
	CreateUsersController(rw http.ResponseWriter, r *http.Request)
	GetAllUsersController(rw http.ResponseWriter, r *http.Request)
	GetUsersByIdController(rw http.ResponseWriter, r *http.Request)
	DeleteUsersByIdController(rw http.ResponseWriter, r *http.Request)
	UpdateUsersByIdController(rw http.ResponseWriter, r *http.Request)
}
