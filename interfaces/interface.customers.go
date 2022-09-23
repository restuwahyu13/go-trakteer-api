package interfaces

import (
	"context"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type ICustomersRepository interface {
	RegisterRepository(ctx context.Context, body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginRepository(ctx context.Context, body *dtos.DTOCustomersLogin) helpers.APIResponse
	ActivationRepository(ctx context.Context, params *dtos.DTOCustomersActivation) helpers.APIResponse
	ResendActivationRepository(ctx context.Context, body *dtos.DTOCustomersResendActivation) helpers.APIResponse
	ForgotPasswordRepository(ctx context.Context, body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordRepository(ctx context.Context, body *dtos.DTOCustomersResetPassword, params *dtos.DTOCustomerResetPasswordToken) helpers.APIResponse
	ChangePasswordRepository(ctx context.Context, body *dtos.DTOCustomersChangePassword, params *dtos.DTOCustomersById) helpers.APIResponse
	GetProfileByIdRepository(ctx context.Context, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdRepository(ctx context.Context, body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
}

type ICustomersService interface {
	RegisterService(ctx context.Context, body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginService(ctx context.Context, body *dtos.DTOCustomersLogin) helpers.APIResponse
	ActivationService(ctx context.Context, params *dtos.DTOCustomersActivation) helpers.APIResponse
	ResendActivationService(ctx context.Context, body *dtos.DTOCustomersResendActivation) helpers.APIResponse
	ForgotPasswordService(ctx context.Context, body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordService(ctx context.Context, body *dtos.DTOCustomersResetPassword, params *dtos.DTOCustomerResetPasswordToken) helpers.APIResponse
	ChangePasswordService(ctx context.Context, body *dtos.DTOCustomersChangePassword, params *dtos.DTOCustomersById) helpers.APIResponse
	GetProfileByIdService(ctx context.Context, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdService(ctx context.Context, body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
}

type ICustomersController interface {
	RegisterController(rw http.ResponseWriter, r *http.Request)
	LoginController(rw http.ResponseWriter, r *http.Request)
	ActivationController(rw http.ResponseWriter, r *http.Request)
	ResendActivationController(rw http.ResponseWriter, r *http.Request)
	ForgotPasswordController(rw http.ResponseWriter, r *http.Request)
	ResetPasswordController(rw http.ResponseWriter, r *http.Request)
	ChangePasswordController(rw http.ResponseWriter, r *http.Request)
	GetProfileByIdController(rw http.ResponseWriter, r *http.Request)
	UpdateProfileByIdController(rw http.ResponseWriter, r *http.Request)
}
