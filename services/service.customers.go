package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type CustomersService struct {
	repository *repositorys.CustomersRepository
}

func NewCustomersService(repository *repositorys.CustomersRepository) *CustomersService {
	return &CustomersService{repository: repository}
}

func (ctx *CustomersService) RegisterService(body *dtos.DTOCustomersRegister) helpers.APIResponse {
	return ctx.repository.RegisterRepository(body)
}

func (ctx *CustomersService) LoginService(body *dtos.DTOCustomersLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(body)
}

func (ctx *CustomersService) ActivationService(params *dtos.DTOCustomersActivation) helpers.APIResponse {
	return ctx.repository.ActivationRepository(params)
}

func (ctx *CustomersService) ResendActivationService(body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	return ctx.repository.ResendActivationRepository(body)
}

func (ctx *CustomersService) ForgotPasswordService(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	return ctx.repository.ForgotPasswordRepository(body)
}

func (ctx *CustomersService) ResetPasswordService(body *dtos.DTOCustomersResetPassword) helpers.APIResponse {
	return ctx.repository.ResetPasswordRepository(body)
}

func (ctx *CustomersService) ChangePasswordService(body *dtos.DTOCustomersChangePassword) helpers.APIResponse {
	return ctx.repository.ChangePasswordRepository(body)
}

func (ctx *CustomersService) GetProfileByIdService(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return ctx.repository.GetProfileByIdRepository(params)
}

func (ctx *CustomersService) UpdateProfileByIdService(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return ctx.repository.UpdateProfileByIdRepository(body, params)
}
