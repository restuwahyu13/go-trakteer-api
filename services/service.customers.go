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

func (ctx *CustomersService) RegisterService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.RegisterRepository(payload)
}

func (ctx *CustomersService) LoginService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(payload)
}

func (ctx *CustomersService) ActivationService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ActivationRepository(payload)
}

func (ctx *CustomersService) ResendActivationService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ResendActivationRepository(payload)
}

func (ctx *CustomersService) ForgotPasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ActivationRepository(payload)
}

func (ctx *CustomersService) ResetPasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ResetPasswordRepository(payload)
}

func (ctx *CustomersService) ChangePasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ChangePasswordRepository(payload)
}

func (ctx *CustomersService) GetProfileService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetProfileRepository(payload)
}

func (ctx *CustomersService) UpdateProfileService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.UpdateProfileRepository(payload)
}
