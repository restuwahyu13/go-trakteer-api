package repositorys

import (
	"context"
	"encoding/json"
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

func (r *customersRepository) RegisterRepository(ctx context.Context, body *dtos.DTOCustomersRegister) helpers.APIResponse {
	customers := models.Customers{}
	roles := models.Roles{}
	catogories := models.Categories{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkUserEmailChan := make(chan error)
	checkRoleIdChan := make(chan error)
	checkCategorieIdChan := make(chan error)

	socialLink := dtos.DTOCustomersSocialLink{
		GithubLink:    "",
		DribbleLink:   "",
		YoutubeLink:   "",
		FacebookLink:  "",
		InstagramLink: "",
		LinkendinLink: "",
	}
	stringify, _ := json.Marshal(socialLink)

	customers.Name = body.Name
	customers.Username = body.Username
	customers.Email = body.Email
	customers.Password = packages.HashPassword(body.Password)
	customers.Active = true
	customers.Verified = false
	customers.SocialLink = json.RawMessage(stringify)
	customers.RoleId = *body.RoleId
	customers.CategorieId = *body.CategorieId

	go (func() {
		checkUserEmail := r.db.GetContext(ctx, &customers, "SELECT username, email FROM customers WHERE username = $1 OR email = $2", customers.Username, customers.Email)
		checkUserEmailChan <- checkUserEmail

		checkRoleId := r.db.GetContext(ctx, &roles, "SELECT id, name FROM roles WHERE id = $1", customers.RoleId)
		checkRoleIdChan <- checkRoleId

		checkCategorieId := r.db.GetContext(ctx, &catogories, "SELECT id FROM categories WHERE id = $1", customers.CategorieId)
		checkCategorieIdChan <- checkCategorieId
	})()

	if <-checkUserEmailChan == nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = fmt.Sprintf("Username %v or Email %v already taken", customers.Username, customers.Email)
		return res
	}

	if err := <-checkRoleIdChan; err != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = fmt.Sprintf("Role data for this id %d not exist", customers.RoleId)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if err := <-checkCategorieIdChan; err != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = fmt.Sprintf("Categorie data for this id %d not exist", customers.CategorieId)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if roles.Name != "customer" {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Your role must be customer"
		return res
	}

	_, insertUserErr := r.db.NamedQueryContext(ctx, `
		INSERT INTO customers (name, username, email, password, active, verified, social_link, role_id, categorie_id)
		VALUES(:name, :username, :email, :password, :active, :verified, :social_link, :role_id, :categorie_id) RETURNING id, email`, &customers)

	if insertUserErr != nil {
		res.StatCode = http.StatusConflict
		res.StatMsg = "Create new customer account failed"
		defer logrus.Errorf("Error Logs: %v", insertUserErr)
		return res
	}

	token.ResourceId = customers.Id
	token.ResourceType = "activation"
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

	htmlTemplateErrchan := make(chan error)
	sendEmailErrChan := make(chan error)

	go func() {
		htmlContent := helpers.HtmlContent{}
		htmlContent.Url = viper.GetString("FE_URL")
		htmlContent.To = customers.Email
		htmlContent.Token = token.AccessToken

		htmlTemplateRes, htmlTemplateErr := helpers.HtmlRender("template.activationAccount", htmlContent)
		htmlTemplateErrchan <- htmlTemplateErr

		sendEmailErr := helpers.SmtpEmail([]string{customers.Email}, "Activation Account!", htmlTemplateRes)
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

	res.StatCode = http.StatusCreated
	res.StatMsg = "Create new customer account success"
	return res
}

/**
* @description LoginRepository
**/

func (r *customersRepository) LoginRepository(ctx context.Context, body *dtos.DTOCustomersLogin) helpers.APIResponse {
	customers := models.Customers{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	customers.Email = body.Email
	customers.Password = body.Password

	checkUserEmail, err := r.db.QueryContext(ctx, `SELECT
		customers.id, customers.name, customers.email, customers.password, customers.verified,
		roles.id as role_id, roles.name as role_name
		FROM customers
		INNER JOIN roles ON customers.role_id = roles.id
		WHERE customers.email = $1
	`, body.Email)

	if err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Csutomer email %v not registered", customers.Email)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if err := carta.Map(checkUserEmail, &customers); err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Relation between table error: %v", err)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if customers.Verified != true {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Customer account for %s not verified, please check your email", body.Email)
		return res
	}

	comparePassword := packages.ComparePassword(body.Password, customers.Password)
	if !comparePassword {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Csutomer password %s miss match", body.Password)
		return res
	}

	jwtPayload := make(map[string]interface{})
	jwtPayload["email"] = customers.Email
	jwtPayload["role"] = customers.Role.Name

	jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	timeFormat := time.RFC1123Z

	accessTokenExpired := helpers.ExpiredAt(1, "days")
	refrehTokenExpired := helpers.ExpiredAt(2, "months")
	expiredAt := time.Now().Add(accessTokenExpired).In(jakartaTimeZone)

	accessToken := packages.SignToken(jwtPayload, accessTokenExpired)
	refrehToken := packages.SignToken(jwtPayload, refrehTokenExpired)

	token.ResourceId = customers.Id
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
		AccessTokenExpired:  time.Now().Add(accessTokenExpired).In(jakartaTimeZone).Format(timeFormat),
		RefreshTokenExpired: time.Now().Add(refrehTokenExpired).In(jakartaTimeZone).Format(timeFormat),
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Login success"
	res.Data = accessTokenPayload
	return res
}

/**
* @description ActivationRepository
**/

func (r *customersRepository) ActivationRepository(ctx context.Context, params *dtos.DTOCustomersActivation) helpers.APIResponse {
	customers := models.Customers{}
	token := models.Token{}
	res := helpers.APIResponse{}

	checkAccessTokenErr := r.db.GetContext(ctx, &token, "SELECT resource_id, expired_at FROM token WHERE access_token = $1 AND resource_type = $2", params.Token, "activation")
	if checkAccessTokenErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Access token not match"
		defer logrus.Errorf("Error Logs: %s", checkAccessTokenErr)
		return res
	}

	jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	timeFormat := "2006-01-02 15:04:05"
	timeNow := time.Now().In(jakartaTimeZone).Format(timeFormat)

	if token.ExpiredAt.Format(timeFormat) < timeNow {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Access token expired, please resend new activation token"
		return res
	}

	customers.Id = token.ResourceId
	customers.Verified = true

	_, updateVerifiedError := r.db.NamedQueryContext(ctx, "UPDATE customers SET verified = :verified WHERE id = :id", &customers)
	if updateVerifiedError != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Verified customer account failed"
		defer logrus.Errorf("Error Logs: %s", updateVerifiedError)
		return res
	}

	res.StatCode = http.StatusBadRequest
	res.StatMsg = "Verified customer account success"
	return res
}

/**
* @description ResendActivationRepository
**/

func (r *customersRepository) ResendActivationRepository(ctx context.Context, body *dtos.DTOCustomersResendActivation) helpers.APIResponse {
	customers := models.Customers{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	customers.Email = body.Email
	checkUserEmailErr := r.db.GetContext(ctx, &customers, "SELECT id, email FROM customers WHERE email = $1", customers.Email)

	if checkUserEmailErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User email %s not exist", customers.Email)
		defer logrus.Errorf("Error Logs: %v", checkUserEmailErr)
		return res
	}

	htmlTemplateErrchan := make(chan error)
	sendEmailErrChan := make(chan error)
	randomToken := helpers.RandomToken()

	go func() {
		htmlContent := helpers.HtmlContent{}
		htmlContent.Url = viper.GetString("FE_URL")
		htmlContent.To = customers.Email
		htmlContent.Token = randomToken

		htmlTemplateRes, htmlTemplateErr := helpers.HtmlRender("template.activationAccount", htmlContent)
		htmlTemplateErrchan <- htmlTemplateErr

		sendEmailErr := helpers.SmtpEmail([]string{customers.Email}, "Activation Account!", htmlTemplateRes)
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

	token.ResourceId = customers.Id
	token.ResourceType = "activation"
	token.AccessToken = randomToken
	token.ExpiredAt = time.Now().Add(helpers.ExpiredAt(5, "minute")).Local()

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
	res.StatMsg = fmt.Sprintf("Resend activation success, please check your email %s address", customers.Email)
	return res
}

/**
* @description ForgotPasswordRepository
**/

func (r *customersRepository) ForgotPasswordRepository(ctx context.Context, body *dtos.DTOCustomersForgotPassword) helpers.APIResponse {
	customers := models.Customers{}
	token := models.Token{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	customers.Email = body.Email
	checkUserEmailErr := r.db.GetContext(ctx, &customers, "SELECT id, email FROM customers WHERE email = $1", customers.Email)

	if checkUserEmailErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("User email %s not exist", customers.Email)
		defer logrus.Errorf("Error Logs: %v", checkUserEmailErr)
		return res
	}

	htmlTemplateErrchan := make(chan error)
	sendEmailErrChan := make(chan error)
	randomToken := helpers.RandomToken()

	go func() {
		htmlContent := helpers.HtmlContent{}
		htmlContent.Url = viper.GetString("FE_URL")
		htmlContent.To = customers.Email
		htmlContent.Token = randomToken

		htmlTemplateRes, htmlTemplateErr := helpers.HtmlRender("template.resetPassword", htmlContent)
		htmlTemplateErrchan <- htmlTemplateErr

		sendEmailErr := helpers.SmtpEmail([]string{customers.Email}, "Reset Password!", htmlTemplateRes)
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

	token.ResourceId = customers.Id
	token.ResourceType = "reset password"
	token.AccessToken = randomToken
	token.ExpiredAt = time.Now().Add(helpers.ExpiredAt(5, "minute")).Local()

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
	res.StatMsg = fmt.Sprintf("Reset password success, please check your email %s address", customers.Email)
	return res
}

/**
* @description ResetPasswordRepository
**/

func (r *customersRepository) ResetPasswordRepository(ctx context.Context, body *dtos.DTOCustomersResetPassword, params *dtos.DTOCustomerResetPasswordToken) helpers.APIResponse {
	customers := models.Customers{}
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

	customers.Id = token.ResourceId
	customers.Password = packages.HashPassword(body.Password)

	_, updatePasswordErr := r.db.NamedQueryContext(ctx, "UPDATE customers SET password = :password WHERE id = :id", &customers)
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

func (r *customersRepository) ChangePasswordRepository(ctx context.Context, body *dtos.DTOCustomersChangePassword, params *dtos.DTOCustomersById) helpers.APIResponse {
	customers := models.Customers{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	if body.Cpassword != body.Password {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Confirm password not match with password"
		return res
	}

	customers.Id = params.Id
	customers.Password = packages.HashPassword(body.Password)

	_, updatePasswordErr := r.db.NamedQueryContext(ctx, "UPDATE customers SET password = :password WHERE id = :id", &customers)
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

func (r *customersRepository) GetProfileByIdRepository(ctx context.Context, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	customers := models.Customers{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	customers.Id = params.Id
	checkUserErr := r.db.GetContext(ctx, &customers, "SELECT id, name, username, email, active, verified, created_at, updated_at, deleted_at FROM customers WHERE id = $1", customers.Id)

	if checkUserErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("UserId for this id %d not exist ", customers.Id)
		defer logrus.Errorf("Error Logs: %v", checkUserErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Get profile data success"
	res.Data = customers
	return res
}

/**
* @description UpdateProfileByIdRepository
**/

func (r *customersRepository) UpdateProfileByIdRepository(ctx context.Context, body *dtos.DTOCustomersUpdateProfileById, params *dtos.DTOCustomersGetProfileById) helpers.APIResponse {
	customers := models.Customers{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkUserErr := r.db.GetContext(ctx, &customers, "SELECT id FROM customers WHERE id = $1", params.Id)
	if checkUserErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("UserId for this id %d not exist ", customers.Id)
		defer logrus.Errorf("Error Logs: %v", checkUserErr)
		return res
	}

	customers.Id = params.Id
	customers.Name = body.Name
	customers.Username = body.Username
	customers.Email = body.Email
	customers.Active = *body.Active

	_, updateProfileErr := r.db.NamedQueryContext(ctx, "UPDATE customers SET name = :name, username = :username, email = :email, active = :active WHERE id = :id", &customers)
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
