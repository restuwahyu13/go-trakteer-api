package repositorys

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/models"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (ctx *UsersRepository) RegisterRepository(body dtos.DTORegister) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	users.Name = body.Name
	users.Username = body.Username
	users.Email = body.Email
	users.Password = packages.HashPassword(body.Password)
	users.RoleId = body.RoleId
	users.CategorieId = body.CategorieId

	checkUserEmail := ctx.db.Get(&users, "SELECT email FROM users WHERE email = $1", users.Email)
	if checkUserEmail == nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = helpers.Strings("Users %s email already taken", users.Email)
	}

	_, err := ctx.db.NamedQuery("INSERT INTO users (name, username, email, password, role_id, categorie_id) VALUES()", users)

	if err != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Create new user account failed"
	}

	res.StatCode = http.StatusCreated
	res.StatMsg = "Create new user account success"
	return res
}

func (ctx *UsersRepository) LoginRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusBadRequest,
		StatMsg:  "Respon from login repository",
	}

	return res
}

func (ctx *UsersRepository) ActivationRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from activation repository",
	}

	return res
}

func (ctx *UsersRepository) ForgotPasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from forgot password repository",
	}

	return res
}

func (ctx *UsersRepository) ResetPasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from reset password repository",
	}

	return res
}

func (ctx *UsersRepository) ChangePasswordRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from change password repository",
	}

	return res
}

func (ctx *UsersRepository) GetProfileRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get profile repository",
	}

	return res
}

func (ctx *UsersRepository) UpdateProfileRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from update profile repository",
	}

	return res
}
