package repositorys

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/models"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type CustomersRepository = interfaces.ICustomersRepository
type customersRepository struct {
	db *sqlx.DB
}

func NewCustomersRepository(db *sqlx.DB) *customersRepository {
	return &customersRepository{db: db}
}

/**
* @description RegisterRepository
**/

func (ctx *customersRepository) RegisterRepository(body *dtos.DTOCustomersRegister) helpers.APIResponse {
	users := models.Users{}
	roles := models.Roles{}
	catogories := models.Categories{}
	res := helpers.APIResponse{}

	checkUserEmailChan := make(chan error)
	checkRoleIdChan := make(chan error)
	checkCategorieIdChan := make(chan error)

	users.Name = body.Name
	users.Username = body.Username
	users.Email = body.Email
	users.Password = packages.HashPassword(body.Password)
	users.Active = true
	users.Verified = false
	users.RoleId = body.RoleId
	users.CategorieId = body.CategorieId

	go (func() {
		checkUserEmail := ctx.db.Get(&users, "SELECT username, email FROM users WHERE username = $1 OR email = $2", users.Username, users.Email)
		checkUserEmailChan <- checkUserEmail

		checkRoleId := ctx.db.Get(&roles, "SELECT id FROM roles WHERE id = $1", body.RoleId)
		checkRoleIdChan <- checkRoleId

		checkCategorieId := ctx.db.Get(&catogories, "SELECT id FROM catogories WHERE id = $1", users.CategorieId)
		checkCategorieIdChan <- checkCategorieId
	})()

	if <-checkUserEmailChan == nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = fmt.Sprintf("Username %v or Email %v already taken", users.Username, users.Email)
		return res
	}

	if <-checkRoleIdChan != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Role name is not exist"
		res.QueryError = <-checkRoleIdChan
		return res
	}

	if <-checkCategorieIdChan != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Categorie name is not exist"
		res.QueryError = <-checkCategorieIdChan
		return res
	}

	if users.Role.Name != "customer" {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Your role must be customer"
		return res
	}

	_, err := ctx.db.NamedQuery(`
		INSERT INTO users (name, username, email, password, active, verified, social_link, role_id, categorie_id)
		VALUES(:name, :username, :email, :password, :active, :verified, :social_link, :role_id, :categorie_id)`, users)

	if err != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Create new customer account failed"
		res.QueryError = err
	}

	res.StatCode = http.StatusCreated
	res.StatMsg = "Create new customer account success"
	return res
}

func (ctx *customersRepository) LoginRepository(body *dtos.DTOCustomersLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from login repository",
	}

	return res
}

func (ctx *customersRepository) ActivationRepository(body *dtos.DTOCustomersActivation) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from activation repository",
	}

	return res
}

func (ctx *customersRepository) ResendActivationRepository(body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from resend activation repository",
	}

	return res
}

func (ctx *customersRepository) ForgotPasswordRepository(body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from forgot password repository",
	}

	return res
}

func (ctx *customersRepository) ResetPasswordRepository(body *dtos.DTOCustomersResetPassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from reset password repository",
	}

	return res
}

func (ctx *customersRepository) ChangePasswordRepository(body *dtos.DTOCustomersChangePassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from change password repository",
	}

	return res
}

func (ctx *customersRepository) GetProfileByIdRepository(params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get profile repository",
	}

	return res
}

func (ctx *customersRepository) UpdateProfileByIdRepository(body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from update profile repository",
	}

	return res
}
