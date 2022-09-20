package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type IUsersRepository interface {
	LoginRepository(ctx context.Context, body *dtos.DTOUsersLogin) helpers.APIResponse
	ForgotPasswordRepository(ctx context.Context, body *dtos.DTOUsersForgotPassword) helpers.APIResponse
	ResetPasswordRepository(ctx context.Context, body *dtos.DTOUsersResetPassword, params *dtos.DTOUsersResetPasswordToken) helpers.APIResponse
	ChangePasswordRepository(ctx context.Context, body *dtos.DTOUsersChangePassword, params *dtos.DTOUsersById) helpers.APIResponse
	GetProfileByIdRepository(ctx context.Context, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	UpdateProfileByIdRepository(ctx context.Context, body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	CreateUsersRepository(ctx context.Context, body *dtos.DTOUsersCreate) helpers.APIResponse
	GetAllUsersRepository(ctx context.Context, query *dtos.DTOUsersPagination) helpers.APIResponse
	GetUsersByIdRepository(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse
	DeleteUsersByIdRepository(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse
	UpdateUsersByIdRepository(ctx context.Context, body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse
}

type IUsersService interface {
	LoginService(ctx context.Context, body *dtos.DTOUsersLogin) helpers.APIResponse
	ForgotPasswordService(ctx context.Context, body *dtos.DTOUsersForgotPassword) helpers.APIResponse
	ResetPasswordService(ctx context.Context, body *dtos.DTOUsersResetPassword, params *dtos.DTOUsersResetPasswordToken) helpers.APIResponse
	ChangePasswordService(ctx context.Context, body *dtos.DTOUsersChangePassword, params *dtos.DTOUsersById) helpers.APIResponse
	GetProfileByIdService(ctx context.Context, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	UpdateProfileByIdService(ctx context.Context, body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse
	CreateUsersService(ctx context.Context, body *dtos.DTOUsersCreate) helpers.APIResponse
	GetAllUsersService(ctx context.Context, query *dtos.DTOUsersPagination) helpers.APIResponse
	GetUsersByIdService(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse
	DeleteUsersByIdService(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse
	UpdateUsersByIdService(ctx context.Context, body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse
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
