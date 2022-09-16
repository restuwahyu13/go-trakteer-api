package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type UsersService = interfaces.IUsersService
type usersService struct {
	repository repositorys.UsersRepository
}

func NewUsersService(repository repositorys.UsersRepository) *usersService {
	return &usersService{repository: repository}
}

func (ctx *usersService) LoginService(body *dtos.DTOUsersLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(body)
}

func (ctx *usersService) ForgotPasswordService(body *dtos.DTOUsersForgotPassword) helpers.APIResponse {
	return ctx.repository.ForgotPasswordRepository(body)
}

func (ctx *usersService) ResetPasswordService(body *dtos.DTOUsersResetPassword) helpers.APIResponse {
	return ctx.repository.ResetPasswordRepository(body)
}

func (ctx *usersService) ChangePasswordService(body *dtos.DTOUsersChangePassword) helpers.APIResponse {
	return ctx.repository.ChangePasswordRepository(body)
}

func (ctx *usersService) GetProfileByIdService(params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	return ctx.repository.GetProfileByIdRepository(params)
}

func (ctx *usersService) UpdateProfileByIdService(body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	return ctx.repository.UpdateProfileByIdRepository(body, params)
}

func (ctx *usersService) CreateUsersService(body *dtos.DTOUsersCreate) helpers.APIResponse {
	return ctx.repository.CreateUsersRepository(body)
}

func (ctx *usersService) GetAllUsersService(query *dtos.DTOUsersPagination) helpers.APIResponse {
	return ctx.repository.GetAllUsersRepository(query)
}

func (ctx *usersService) GetUsersByIdService(params *dtos.DTOUsersById) helpers.APIResponse {
	return ctx.repository.GetUsersByIdRepository(params)
}

func (ctx *usersService) DeleteUsersByIdService(params *dtos.DTOUsersById) helpers.APIResponse {
	return ctx.repository.DeleteUsersByIdRepository(params)
}

func (ctx *usersService) UpdateUsersByIdService(body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse {
	return ctx.repository.UpdateUsersByIdRepository(body, params)
}
