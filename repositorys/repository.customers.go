package repositorys

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type CustomersRepository struct {
	db *sqlx.DB
}

func NewCustomersRepository(db *sqlx.DB) *CustomersRepository {
	return &CustomersRepository{db: db}
}

func (ctx *CustomersRepository) RegisterRepository(body *dtos.DTOCustomersRegister) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from register repository",
	}

	return res
}

func (ctx *CustomersRepository) LoginRepository(body *dtos.DTOCustomersLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from login repository",
	}

	return res
}

func (ctx *CustomersRepository) ActivationRepository(body *dtos.DTOCustomersActivation) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from activation repository",
	}

	return res
}

func (ctx *CustomersRepository) ResendActivationRepository(body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from resend activation repository",
	}

	return res
}

func (ctx *CustomersRepository) ForgotPasswordRepository(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from forgot password repository",
	}

	return res
}

func (ctx *CustomersRepository) ResetPasswordRepository(body *dtos.DTOCustomersResetPassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from reset password repository",
	}

	return res
}

func (ctx *CustomersRepository) ChangePasswordRepository(body *dtos.DTOCustomersChangePassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from change password repository",
	}

	return res
}

func (ctx *CustomersRepository) GetProfileByIdRepository(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get profile repository",
	}

	return res
}

func (ctx *CustomersRepository) UpdateProfileByIdRepository(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from update profile repository",
	}

	return res
}
