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

func (ctx *CustomersRepository) RegisterRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from register repository",
	}

	return res
}

func (ctx *CustomersRepository) LoginRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from login repository",
	}

	return res
}

func (ctx *CustomersRepository) ActivationRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from activation repository",
	}

	return res
}

func (ctx *CustomersRepository) ResendActivationRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from resend activation repository",
	}

	return res
}

func (ctx *CustomersRepository) ForgotPasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from forgot password repository",
	}

	return res
}

func (ctx *CustomersRepository) ResetPasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from reset password repository",
	}

	return res
}

func (ctx *CustomersRepository) ChangePasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from change password repository",
	}

	return res
}

func (ctx *CustomersRepository) GetProfileRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get profile repository",
	}

	return res
}

func (ctx *CustomersRepository) UpdateProfileRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from update profile repository",
	}

	return res
}
