package services

import (
	"context"

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

func (s *customersService) RegisterService(ctx context.Context, body *dtos.DTOCustomersRegister) helpers.APIResponse {
	return s.repository.RegisterRepository(ctx, body)
}

func (s *customersService) LoginService(ctx context.Context, body *dtos.DTOCustomersLogin) helpers.APIResponse {
	return s.repository.LoginRepository(ctx, body)
}

func (s *customersService) ActivationService(ctx context.Context, params *dtos.DTOCustomersActivation) helpers.APIResponse {
	return s.repository.ActivationRepository(ctx, params)
}

func (s *customersService) ResendActivationService(ctx context.Context, body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	return s.repository.ResendActivationRepository(ctx, body)
}

func (s *customersService) ForgotPasswordService(ctx context.Context, body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	return s.repository.ForgotPasswordRepository(ctx, body)
}

func (s *customersService) ResetPasswordService(ctx context.Context, body *dtos.DTOCustomersResetPassword, params *dtos.DTOCustomerResetPasswordToken) helpers.APIResponse {
	return s.repository.ResetPasswordRepository(ctx, body, params)
}

func (s *customersService) ChangePasswordService(ctx context.Context, body *dtos.DTOCustomersChangePassword, params *dtos.DTOCustomersById) helpers.APIResponse {
	return s.repository.ChangePasswordRepository(ctx, body, params)
}

func (s *customersService) GetProfileByIdService(ctx context.Context, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return s.repository.GetProfileByIdRepository(ctx, params)
}

func (s *customersService) UpdateProfileByIdService(ctx context.Context, body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	return s.repository.UpdateProfileByIdRepository(ctx, body, params)
}
