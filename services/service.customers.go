package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type CustomersService = interfaces.ICustomersService
type customersService struct {
	repository repositorys.CustomersRepository
}

func NewCustomersService(repository repositorys.CustomersRepository) *customersService {
	return &customersService{repository: repository}
}

func (ctx *customersService) RegisterService(body *dtos.DTOCustomersRegister) helpers.APIResponse {
	return ctx.repository.RegisterRepository(body)
}

func (ctx *customersService) LoginService(body *dtos.DTOCustomersLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(body)
}

func (ctx *customersService) ActivationService(params *dtos.DTOCustomersActivation) helpers.APIResponse {
	return ctx.repository.ActivationRepository(params)
}

func (ctx *customersService) ResendActivationService(body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	return ctx.repository.ResendActivationRepository(body)
}

func (ctx *customersService) ForgotPasswordService(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	return ctx.repository.ForgotPasswordRepository(body)
}

func (ctx *customersService) ResetPasswordService(body *dtos.DTOCustomersResetPassword) helpers.APIResponse {
	return ctx.repository.ResetPasswordRepository(body)
}

func (ctx *customersService) ChangePasswordService(body *dtos.DTOCustomersChangePassword) helpers.APIResponse {
	return ctx.repository.ChangePasswordRepository(body)
}

func (ctx *customersService) GetProfileByIdService(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return ctx.repository.GetProfileByIdRepository(params)
}

func (ctx *customersService) UpdateProfileByIdService(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return ctx.repository.UpdateProfileByIdRepository(body, params)
}
