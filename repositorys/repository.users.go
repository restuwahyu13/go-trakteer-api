package repositorys

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

var (
	min = 10 * time.Second
	max = 60 * time.Second
)

func NewUsersRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{db: db}
}

/**
* @description LoginRepository
**/

func (r *usersRepository) LoginRepository(ctx context.Context, body *dtos.DTOUsersLogin) helpers.APIResponse {
	users := models.Users{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Email = body.Email
	users.Password = body.Password

	checkUserEmail, err := r.db.QueryContext(ctx, `
		SELECT users.id, users.name, users.email, users.password, roles.id as role_id, roles.name as role_name
		FROM users INNER JOIN roles ON users.role_id = roles.id WHERE users.email = $1
	`, body.Email)

	if err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Users email %v not registered", users.Email)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if err := carta.Map(checkUserEmail, &users); err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Relation between table error: %v", err)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	comparePassword := packages.ComparePassword(body.Password, users.Password)

	if !comparePassword {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Users password %s miss match", body.Password)
		return res
	}

	jwtPayload := make(map[string]interface{})
	jwtPayload["id"] = users.Id
	jwtPayload["role"] = users.Role.Name

	jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	timeFormat := time.RFC1123Z

	accessTokenExpired := helpers.ExpiredAt(1, "days")
	refrehTokenExpired := helpers.ExpiredAt(2, "months")
	expiredAt := time.Now().Add(accessTokenExpired).In(jakartaTimeZone)

	accessToken := packages.SignToken(jwtPayload, accessTokenExpired)
	refrehToken := packages.SignToken(jwtPayload, refrehTokenExpired)

	token.ResourceId = users.Id
	token.ResourceType = "login"
	token.AccessToken = accessToken
	token.RefreshToken = refrehToken
	token.ExpiredAt = expiredAt

	_, insertTokenErr := r.db.NamedQueryContext(ctx, `
	INSERT INTO token (resource_id, resource_type, access_token, refresh_token, expired_at)
	VALUES (:resource_id, :resource_type, :access_token, :refresh_token, :expired_at)`, &token)

	if insertTokenErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Insert token into database failed"
		defer logrus.Errorf("Error Logs: %v", insertTokenErr)
		return res
	}

	accessTokenPayload := packages.UsersToken{
		AccessToken:         accessToken,
		RefreshToken:        refrehToken,
		AccessTokenExpired:  time.Now().Add(time.Duration(accessTokenExpired)).In(jakartaTimeZone).Format(timeFormat),
		RefreshTokenExpired: time.Now().Add(time.Duration(refrehTokenExpired)).In(jakartaTimeZone).Format(timeFormat),
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Login success"
	res.Data = accessTokenPayload
	return res
}

/**
* @description ForgotPasswordRepository
**/

func (r *usersRepository) ForgotPasswordRepository(ctx context.Context, body *dtos.DTOUsersForgotPassword) helpers.APIResponse {
	users := models.Users{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Email = body.Email
	checkUserEmailErr := r.db.GetContext(ctx, &users, "SELECT id, email FROM users WHERE email = $1", users.Email)

	if checkUserEmailErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User email %s not exist", users.Email)
		defer logrus.Errorf("Error Logs: %v", checkUserEmailErr)
		return res
	}

	htmlTemplateErrchan := make(chan error)
	sendEmailErrChan := make(chan error)

	go func() {
		htmlContent := helpers.HtmlContent{}
		htmlContent.Url = viper.GetString("FE_URL")
		htmlContent.To = users.Email
		htmlContent.Token = helpers.RandomToken()

		htmlTemplateRes, htmlTemplateErr := helpers.HtmlRender("template.resetPassword", htmlContent)
		htmlTemplateErrchan <- htmlTemplateErr

		sendEmailErr := helpers.SmtpEmail([]string{users.Email}, "Reset Password!", htmlTemplateRes)
		sendEmailErrChan <- sendEmailErr
	}()

	if err := <-htmlTemplateErrchan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Render html template error: %v", err)
		defer logrus.Errorf("Error Logs: %v", err)
		defer close(htmlTemplateErrchan)
		return res
	}

	if err := <-sendEmailErrChan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Send smtp email error: %v", err)
		defer logrus.Errorf("Error Logs: %v", err)
		defer close(htmlTemplateErrchan)
		return res
	}

	token.ResourceId = users.Id
	token.ResourceType = "reset password"
	token.AccessToken = helpers.RandomToken()
	token.ExpiredAt = time.Now().Add(time.Duration(helpers.ExpiredAt(5, "minutes")))

	_, insertTokenErr := r.db.NamedQueryContext(ctx, `
	INSERT INTO token (resource_id, resource_type, access_token, refresh_token, expired_at)
	VALUES (:resource_id, :resource_type, :access_token, :refresh_token, :expired_at)`, &token)

	if insertTokenErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Insert token into database failed"
		defer logrus.Errorf("Error Logs: %v", insertTokenErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Reset password success, please check your email %s address", users.Email)
	return res
}

/**
* @description ResetPasswordRepository
**/

func (r *usersRepository) ResetPasswordRepository(ctx context.Context, body *dtos.DTOUsersResetPassword, params *dtos.DTOUsersResetPasswordToken) helpers.APIResponse {
	users := models.Users{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkAccessTokenErr := r.db.GetContext(ctx, &token, "SELECT resource_id, expired_at FROM token WHERE access_token = $1 AND resource_type = $2 ORDER BY id DESC", params.Token, "reset password")
	if checkAccessTokenErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Invalid token format or Token not match"
		defer logrus.Errorf("Error Logs: %v", checkAccessTokenErr)
		return res
	}

	jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	timeFormat := "2006-01-02 15:04:05"
	timeNow := time.Now().In(jakartaTimeZone).Format(timeFormat)

	if token.ExpiredAt.Format(timeFormat) < timeNow {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Token expired, please try forgot password"
	}

	if body.Cpassword != body.Password {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Confirm password not match with password"
		return res
	}

	users.Id = token.ResourceId
	users.Password = packages.HashPassword(body.Password)

	_, updatePasswordErr := r.db.NamedQueryContext(ctx, "UPDATE users SET password = :password WHERE id = :id", &users)
	if updatePasswordErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Update reset password account failed"
		defer logrus.Errorf("Error Logs: %v", updatePasswordErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Reset old password to new password success"
	return res
}

/**
* @description ChangePasswordRepository
**/

func (r *usersRepository) ChangePasswordRepository(ctx context.Context, body *dtos.DTOUsersChangePassword, params *dtos.DTOUsersById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	if body.Cpassword != body.Password {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Confirm password not match with password"
		return res
	}

	users.Id = params.Id
	users.Password = packages.HashPassword(body.Password)

	_, updatePasswordErr := r.db.NamedQueryContext(ctx, "UPDATE users SET password = :password WHERE id = :id", &users)
	if updatePasswordErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Change old password to new password failed"
		defer logrus.Errorf("Error Logs: %v", updatePasswordErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Change old password to new password success"
	return res
}

/**
* @description GetProfileByIdRepository
**/

func (r *usersRepository) GetProfileByIdRepository(ctx context.Context, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Id = params.Id
	checkUserErr := r.db.GetContext(ctx, &users, "SELECT id, name, email, active, verified, created_at, updated_at, deleted_at FROM users WHERE id = $1", users.Id)

	if checkUserErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User data for this id %d not exist ", users.Id)
		defer logrus.Errorf("Error Logs: %v", checkUserErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Get profile data success"
	res.Data = users
	return res
}

/**
* @description UpdateProfileByIdRepository
**/

func (r *usersRepository) UpdateProfileByIdRepository(ctx context.Context, body *dtos.DTOUsersUpdateProfileById, params *dtos.DTOUsersGetProfileById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkUserErr := r.db.GetContext(ctx, &users, "SELECT id FROM users WHERE id = $1", params.Id)
	if checkUserErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User data for this id %d not exist ", users.Id)
		defer logrus.Errorf("Error Logs: %v", checkUserErr)
		return res
	}

	users.Id = params.Id
	users.Name = body.Name
	users.Email = body.Email
	users.Active = *body.Active

	_, updateProfileErr := r.db.NamedQueryContext(ctx, "UPDATE users SET name = :name, email = :email, active = :active WHERE id = :id", &users)
	if updateProfileErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Update profile failed"
		defer logrus.Errorf("Error Logs: %v", updateProfileErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated profile success"
	return res
}

/**
* @description CreateUsersRepository
**/

func (r *usersRepository) CreateUsersRepository(ctx context.Context, body *dtos.DTOUsersCreate) helpers.APIResponse {
	users := models.Users{}
	roles := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Name = body.Name
	users.Email = body.Email
	users.Password = packages.HashPassword(body.Password)
	users.Active = true
	users.Verified = true
	users.RoleId = body.RoleId

	checkUserErrChan := make(chan error)
	checkRoleErrChan := make(chan error)

	go func() {
		checkUserErr := r.db.GetContext(ctx, &users, "SELECT email FROM users WHERE email = $2", users.Email)
		checkUserErrChan <- checkUserErr

		checkRoleErr := r.db.GetContext(ctx, &roles, "SELECT id FROM roles WHERE id = $1", users.RoleId)
		checkRoleErrChan <- checkRoleErr
	}()

	if err := <-checkUserErrChan; err == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Email %s already taken", users.Email)
		defer logrus.Errorf("Error Logs: %v", err)
		defer close(checkUserErrChan)
		return res
	}

	if err := <-checkRoleErrChan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d not exist", users.RoleId)
		defer logrus.Errorf("Error Logs: %v", err)
		defer close(checkRoleErrChan)
		return res
	}

	_, insertErr := r.db.NamedQueryContext(ctx, `
		INSERT INTO users (name, email, password, active, verified, role_id)
		VALUES(:name, :email, :password, :active, :verified, :role_id)`,
		&users)

	if insertErr != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Create new user account failed"
		defer logrus.Errorf("Error Logs: %v", insertErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Create new user success"
	return res
}

/**
* @description GetAllUsersRepository
**/

func (r *usersRepository) GetAllUsersRepository(ctx context.Context, query *dtos.DTOUsersPagination) helpers.APIResponse {
	users := []models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, max)
	defer cancel()

	getAllUsersChan := make(chan error)
	countChan := make(chan int)

	go func() {
		getAllRoles := r.db.SelectContext(ctx, &users, fmt.Sprintf("SELECT id, name, email, active, verified, created_at, updated_at, deleted_at FROM users ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)
		getAllUsersChan <- getAllRoles

		count := 0
		countRoles := r.db.QueryRowContext(ctx, "SELECT COUNT(id) FROM users")
		countRoles.Scan(&count)
		countChan <- count
	}()

	if <-getAllUsersChan != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Users data not exist"
		defer logrus.Errorf("Error Logs: %v", <-getAllUsersChan)
		defer close(getAllUsersChan)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Users data already to use"
	res.Data = users
	res.Pagination = helpers.Stringify(helpers.Pagination(query, <-countChan))
	defer close(countChan)
	return res
}

/**
* @description GetUsersByIdRepository
**/

func (r *usersRepository) GetUsersByIdRepository(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Id = params.Id
	getRoleId := r.db.GetContext(ctx, &users, "SELECT id, name, username, email, active, verified, created_at, updated_at, deleted_at FROM users WHERE id = $1", users.Id)

	if getRoleId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User data for this id %d not exist", users.Id)
		defer logrus.Errorf("Error Logs: %v", getRoleId)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "User data already to use"
	res.Data = users
	return res
}

/**
* @description DeleteUsersByIdRepository
**/

func (r *usersRepository) DeleteUsersByIdRepository(ctx context.Context, params *dtos.DTOUsersById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Id = params.Id
	checkRoleId := r.db.GetContext(ctx, &users, "SELECT id FROM users WHERE id = $1", users.Id)

	if checkRoleId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User data for this id %d not exist", users.Id)
		defer logrus.Errorf("Error Logs: %v", checkRoleId)
		return res
	}

	deletedTime := time.Now().Local()
	users.DeletedAt = &deletedTime

	_, deletedRoleErr := r.db.NamedQueryContext(ctx, "UPDATE users SET deleted_at = :deleted_at WHERE id = :id", &users)

	if deletedRoleErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Deleted user for this id %d failed", users.Id)
		defer logrus.Errorf("Error Logs: %v", deletedRoleErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Deleted user data for this id %d success", users.Id)
	return res
}

/**
* @description UpdateUsersByIdRepository
**/

func (r *usersRepository) UpdateUsersByIdRepository(ctx context.Context, body *dtos.DTOUsersUpdate, params *dtos.DTOUsersById) helpers.APIResponse {
	users := models.Users{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	users.Id = params.Id
	checkUserId := r.db.GetContext(ctx, &users, "SELECT id FROM users WHERE id = $1", users.Id)

	if checkUserId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User data for this id %d not exist", users.Id)
		defer logrus.Errorf("Error Logs: %v", checkUserId)
		return res
	}

	users.Name = body.Name
	users.Email = body.Email
	users.Active = *body.Active
	users.RoleId = body.RoleId
	users.UpdatedAt = time.Now()

	_, updatedRoleErr := r.db.NamedQueryContext(ctx, "UPDATE users SET name = :name, username = :username, email = :email, active = :active WHERE id = :id", &users)

	if updatedRoleErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Updated user data for this id %v failed", users.Id)
		defer logrus.Errorf("Error Logs: %v", updatedRoleErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Updated user data for this id %v success", users.Id)
	return res
}
