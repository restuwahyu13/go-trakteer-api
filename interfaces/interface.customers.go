package interfaces

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type ICustomersRepository interface {
	RegisterRepository(body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginRepository(body *dtos.DTOCustomersLogin) helpers.APIResponse
	AcitvationRepository(params *dtos.DTOCustomersActivation) helpers.APIResponse
	ForgotPasswordRepository(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordRepository(body *dtos.DTOCustomersResetPassword) helpers.APIResponse
	ChangePasswordRepository(body *dtos.DTOCustomersChangePassword) helpers.APIResponse
	GetProfileByIdRepository(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdRepository(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
}

type ICustomersService interface {
	RegisterService(body *dtos.DTOCustomersRegister) helpers.APIResponse
	LoginService(body *dtos.DTOCustomersLogin) helpers.APIResponse
	AcitvationService(params *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ForgotPasswordService(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse
	ResetPasswordService(body *dtos.DTOCustomersResetPassword) helpers.APIResponse
	ChangePasswordService(body *dtos.DTOCustomersChangePassword) helpers.APIResponse
	GetProfileByIdService(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
	UpdateProfileByIdService(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse
}

type ICustomersController interface {
	RegisterService(rw http.ResponseWriter, r *http.Request)
	LoginService(rw http.ResponseWriter, r *http.Request)
	AcitvationService(rw http.ResponseWriter, r *http.Request)
	ForgotPasswordService(rw http.ResponseWriter, r *http.Request)
	ResetPasswordService(rw http.ResponseWriter, r *http.Request)
	ChangePasswordService(rw http.ResponseWriter, r *http.Request)
	GetProfileByIdService(rw http.ResponseWriter, r *http.Request)
	UpdateProfileByIdService(rw http.ResponseWriter, r *http.Request)
	GetAllCustomersController(rw http.ResponseWriter, r *http.Request)
}
