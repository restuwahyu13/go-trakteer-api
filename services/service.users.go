package services

import (
	"context"

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

func (s *usersService) LoginService(ctx context.Context, body *dtos.DTOUsersLogin) helpers.APIResponse {
	return s.repository.LoginRepository(ctx, body)
}

func (s *usersService) ForgotPasswordService(ctx context.Context, body *dtos.DTOUsersForgotPassword) helpers.APIResponse {
	return s.repository.ForgotPasswordRepository(ctx, body)
}

func (s *usersService) ResetPasswordService(ctx context.Context, body *dtos.DTOUsersResetPassword, params *dtos.DTOUsersResetPasswordToken) helpers.APIResponse {
	return s.repository.ResetPasswordRepository(ctx, body, params)
}

func (s *usersService) ChangePasswordService(ctx context.Context, body *dtos.DTOUsersChangePassword, params *dtos.DTOUsersById) helpers.APIResponse {
	return s.repository.ChangePasswordRepository(ctx, body, params)
}

func (s *usersService) GetProfileByIdService(ctx context.Context, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	return s.repository.GetProfileByIdRepository(ctx, params)
}

func (s *usersService) UpdateProfileByIdService(ctx context.Context, body *dtos.DTOUsersProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	return s.repository.UpdateProfileByIdRepository(ctx, body, params)
}

func (s *usersService) CreateUsersService(ctx context.Context, body *dtos.DTOUsers) helpers.APIResponse {
	return s.repository.CreateUsersRepository(ctx, body)
}

func (s *usersService) GetAllUsersService(ctx context.Context, query *dtos.DTOUsersPagination) helpers.APIResponse {
	return s.repository.GetAllUsersRepository(ctx, query)
}

func (s *usersService) GetUsersByIdService(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse {
	return s.repository.GetUsersByIdRepository(ctx, params)
}

func (s *usersService) DeleteUsersByIdService(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse {
	return s.repository.DeleteUsersByIdRepository(ctx, params)
}

func (s *usersService) UpdateUsersByIdService(ctx context.Context, body *dtos.DTOUsers, params *dtos.DTOUsersById) helpers.APIResponse {
	return s.repository.UpdateUsersByIdRepository(ctx, body, params)
}

func (s *usersService) HealthCheckTokenService(ctx context.Context, params *dtos.DTOUsersHealthToken) helpers.APIResponse {
	return s.repository.HealthCheckTokenRepository(ctx, params)
}

func (s *usersService) RefreshTokenService(ctx context.Context, body *dtos.DTOUsersRefreshToken) helpers.APIResponse {
	return s.repository.RefreshTokenRepository(ctx, body)
}
