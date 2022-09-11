package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type UsersService struct {
	repository *repositorys.UsersRepository
}

func NewUsersService(repository *repositorys.UsersRepository) *UsersService {
	return &UsersService{repository: repository}
}

func (ctx *UsersService) RegisterService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.RegisterRepository(payload)
}

func (ctx *UsersService) LoginService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(payload)
}

func (ctx *UsersService) ActivationService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ActivationRepository(payload)
}

func (ctx *UsersService) ForgotPasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ForgotPasswordRepository(payload)
}

func (ctx *UsersService) ResetPasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ResetPasswordRepository(payload)
}

func (ctx *UsersService) ChangePasswordService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.ChangePasswordRepository(payload)
}

func (ctx *UsersService) GetProfileService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.GetProfileRepository(payload)
}

func (ctx *UsersService) UpdateProfileService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.UpdateProfileRepository(payload)
}
