package interfaces

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type ICustomersRepository interface {
	RegisterRepository(body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginRepository(body *dtos.DTOCustomersLogin) helpers.APIResponse
	ActivationRepository(params *dtos.DTOCustomersActivation) helpers.APIResponse
	ResendActivationRepository(body *dtos.DTOCustomersResendActivation) helpers.APIResponse
	ForgotPasswordRepository(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordRepository(body *dtos.DTOCustomersResetPassword) helpers.APIResponse
	ChangePasswordRepository(body *dtos.DTOCustomersChangePassword) helpers.APIResponse
	GetProfileByIdRepository(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdRepository(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
}

type ICustomersService interface {
	RegisterService(body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginService(body *dtos.DTOCustomersLogin) helpers.APIResponse
	ActivationService(params *dtos.DTOCustomersActivation) helpers.APIResponse
	ResendActivationService(body *dtos.DTOCustomersResendActivation) helpers.APIResponse
	ForgotPasswordService(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordService(body *dtos.DTOCustomersResetPassword) helpers.APIResponse
	ChangePasswordService(body *dtos.DTOCustomersChangePassword) helpers.APIResponse
	GetProfileByIdService(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdService(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
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
