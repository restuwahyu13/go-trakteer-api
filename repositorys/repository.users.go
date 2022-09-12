package repositorys

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/models"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type UsersRepository struct {
	db *sqlx.DB
}

type usersToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

/**
* @description RegisterRepository
**/

func (ctx *UsersRepository) RegisterRepository(body dtos.DTORegister) helpers.APIResponse {
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
		return res
	}

	if <-checkCategorieIdChan != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Categorie name is not exist"
		return res
	}

	_, err := ctx.db.NamedQuery(`INSERT INTO users (name, username, email, password, active, verified, social_link, role_id, categorie_id)
		VALUES(:name, :username, :email, :password, :active, :verified, :social_link, :role_id, :categorie_id)`, users)

	if err != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Create new user account failed"
	}

	res.StatCode = http.StatusCreated
	res.StatMsg = "Create new user account success"
	return res
}

/**
* @description LoginRepository
**/

func (ctx *UsersRepository) LoginRepository(body dtos.DTOLogin) helpers.APIResponse {
	users := models.Users{}
	token := models.Token{}
	res := helpers.APIResponse{}

	users.Email = body.Email
	users.Password = body.Password

	checkUserEmail, err := ctx.db.Query(`SELECT
		users.id, users.name, users.email, users.password,
		roles.id as role_id, roles.name as role_name, roles.created_at as role_created_at, roles.updated_at as role_updated_at
		FROM users INNER JOIN roles ON users.role_id = roles.id WHERE users.email = $1
	`, body.Email)
	carta.Map(checkUserEmail, &users)

	if err != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Users email %v not registered", users.Email)
		return res
	}

	compare := packages.ComparePassword(body.Password, users.Password)

	if !compare {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Users password %s miss match", body.Password)
		return res
	}

	jwtPayload := make(map[string]interface{})
	jwtPayload["email"] = users.Email
	jwtPayload["role"] = users.Role.Name

	accessToken := packages.SignToken(jwtPayload, 60)    // 1 days
	refrehToken := packages.SignToken(jwtPayload, 86400) // 2 months
	expiredAt := time.Now().Add(time.Duration(time.Minute * 60)).Local()

	token.ResourceID = users.ID
	token.ResourceBy = "login"
	token.AccessToken = accessToken
	token.RefreshToken = refrehToken
	token.ExpiredAt = expiredAt

	accessTokenPayload := usersToken{
		AccessToken:  accessToken,
		RefreshToken: refrehToken,
		Role:         users.Role.Name,
	}

	ctx.db.NamedQuery(`INSERT INTO token (resource_id, resource_by, access_token, refresh_token, expired_at)
	VALUES (:resource_id, :resource_by. :access_token, :refresh_token, :expired_At)`, &token)

	res.StatCode = http.StatusOK
	res.StatMsg = "Login successfully"
	res.Data = accessTokenPayload
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
