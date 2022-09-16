package repositorys

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/models"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type UsersRepository = interfaces.IUsersRepository
type usersRepository struct {
	db *sqlx.DB
}

type usersRole struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type usersToken struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenExpired  string `json:"access_token_expired"`
	RefreshTokenExpired string `json:"refresh_token_expired"`
	User                usersRole
}

func NewUsersRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{db: db}
}

/**
* @description LoginRepository
**/

func (ctx *usersRepository) LoginRepository(body *dtos.DTOUsersLogin) helpers.APIResponse {
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
		res.QueryError = err
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

	accessTokenExpired := helpers.ExpiredAt(1, "days")
	refrehTokenExpired := helpers.ExpiredAt(2, "months")
	jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	expiredAt := time.Now().Add(time.Duration(accessTokenExpired)).In(jakartaTimeZone)

	accessToken := packages.SignToken(jwtPayload, time.Duration(accessTokenExpired))
	refrehToken := packages.SignToken(jwtPayload, time.Duration(refrehTokenExpired))

	token.ResourceId = users.ID
	token.ResourceType = "login"
	token.AccessToken = accessToken
	token.RefreshToken = refrehToken
	token.ExpiredAt = expiredAt

	_, insertTokenErr := ctx.db.NamedQuery(`
	INSERT INTO token (resource_id, resource_type, access_token, refresh_token, expired_at)
	VALUES (:resource_id, :resource_type, :access_token, :refresh_token, :expired_at)`, &token)

	if insertTokenErr != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = "Insert token into database failed"
		res.QueryError = insertTokenErr
		return res
	}

	accessTokenPayload := usersToken{
		AccessToken:         accessToken,
		RefreshToken:        refrehToken,
		AccessTokenExpired:  helpers.TimeFormat(time.Now().Add(time.Duration(accessTokenExpired)).In(jakartaTimeZone)),
		RefreshTokenExpired: helpers.TimeFormat(time.Now().Add(time.Duration(refrehTokenExpired)).In(jakartaTimeZone)),
		User:                usersRole{Name: users.Name, Role: users.Role.Name},
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Login successfully"
	res.Data = accessTokenPayload
	return res
}

func (ctx *usersRepository) ActivationRepository(params *dtos.DTOUsersLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from activation repository",
	}

	return res
}

func (ctx *usersRepository) ForgotPasswordRepository(body *dtos.DTOUsersForgotPassword) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	users.Email = body.Email

	checkUserEmail := ctx.db.Get(&users, "SELECT email FROM users WHERE email = $1", users.Email)
	if checkUserEmail != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("User email %s not exist", users.Email)
		return res
	}

	htmlContent := helpers.HtmlContent{}
	htmlContent.Url = viper.GetString("FE_URL")
	htmlContent.To = users.Email
	htmlContent.Token = helpers.RandomToken()

	htmlTemplateRes, htmlTemplateErr := helpers.HtmlRender("template.resetPassword", htmlContent)

	if htmlTemplateErr != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Render html template error: %v", htmlTemplateErr)
		return res
	}

	sendEmailErr := helpers.SmtpEmail([]string{users.Email}, htmlTemplateRes)

	if sendEmailErr != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Send smtp email error: %v", sendEmailErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Reset password successfully, please check your email %s address", users.Email)
	return res
}

func (ctx *usersRepository) ResetPasswordRepository(body *dtos.DTOUsersResetPassword) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from reset password repository",
	}

	return res
}

func (ctx *usersRepository) ChangePasswordRepository(body *dtos.DTOUsersChangePassword) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	if body.Cpassword != body.Password {
		res.StatCode = http.StatusNotFound
		res.StatMsg = "Confirm password not match with password"
		return res
	}

	users.Password = packages.HashPassword(body.Password)

	_, updatePasswordErr := ctx.db.NamedQuery("UPDATE users SET password = :password WHERE id = :id", &users)
	if updatePasswordErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Change old password to new password failed"
		res.QueryError = updatePasswordErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Change old password to new password successfully"
	return res
}

func (ctx *usersRepository) GetProfileByIdRepository(params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get profile repository",
	}

	return res
}

func (ctx *usersRepository) UpdateProfileByIdRepository(body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from update profile repository",
	}

	return res
}

func (ctx *usersRepository) CreateUsersRepository(body *dtos.DTOUsersCreate) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get all users repository",
	}

	return res
}

func (ctx *usersRepository) GetAllUsersRepository(query *dtos.DTOUsersPagination) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get all users repository",
	}

	return res
}

func (ctx *usersRepository) GetUsersByIdRepository(params *dtos.DTOUsersById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get all users repository",
	}

	return res
}

func (ctx *usersRepository) DeleteUsersByIdRepository(params *dtos.DTOUsersById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get all users repository",
	}

	return res
}

func (ctx *usersRepository) UpdateUsersByIdRepository(body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from get all users repository",
	}

	return res
}
