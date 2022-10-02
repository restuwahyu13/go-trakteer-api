package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/models"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type authHandler struct {
	db *sqlx.DB
}

func NewMiddlewareAuth(db *sqlx.DB) *authHandler {
	return &authHandler{db: db}
}

func (m *authHandler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		res := helpers.APIResponse{}
		headers := r.Header.Get("Authorization")

		if headers == "" {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "Authorization is required"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		bearerToken := strings.Split(headers, " ")

		if bearerToken[0] != "Bearer" {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "Authorization must be using bearer"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		if _, err := packages.VerifyToken(strings.TrimSpace(bearerToken[1])); err != nil {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "AccessToken invalid or expired"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(5*time.Second))
		defer cancel()

		users := models.Users{}
		customers := models.Customers{}
		token := models.Token{}

		decodeToken, _ := packages.ParseToken(bearerToken[1])
		metadataToken := make(map[string]interface{})

		encoded, _ := json.Marshal(&decodeToken)
		json.Unmarshal(encoded, &metadataToken)

		var queryDatabase string
		if metadataToken["role"] == "customer" {
			queryDatabase = `SELECT
				customers.id, customers.email, customers.active,
				roles.id as role_id, roles.name as role_name,
				categories.id as categorie_id, categories.name as categorie_name
				FROM customers
				INNER JOIN roles ON customers.role_id = roles.id
				INNER JOIN categories ON customers.categorie_id = categories.id
				WHERE customers.id = $1 AND customers.active = $2
			`
		} else {
			queryDatabase = `SELECT
				users.id, users.email, users.active,
				roles.id as role_id, roles.name as role_name
				FROM users
				INNER JOIN roles ON users.role_id = roles.id
				WHERE users.id = $1 AND users.active = $2
			`
		}

		checkByIdRows, checkByIdErr := m.db.QueryContext(ctx, queryDatabase, metadataToken["id"], "true")

		if checkByIdErr != nil {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "Metadata accessToken not match with metadataToken from db"
			helpers.Send(rw, helpers.ApiResponse(res))

			defer logrus.Errorf("Error Logs: %v", checkByIdErr)
			return
		}

		if metadataToken["role"] == "customer" {
			relationErr := carta.Map(checkByIdRows, &customers)
			if relationErr != nil {
				defer logrus.Errorf("Error Logs: %v", relationErr)
				return
			}
		} else {
			relationErr := carta.Map(checkByIdRows, &users)
			if relationErr != nil {
				defer logrus.Errorf("Error Logs: %v", relationErr)
				return
			}
		}

		var resourceId int
		if metadataToken["role"] == "customer" {
			resourceId = customers.Id
		} else {
			resourceId = users.Id
		}

		checkTokenErr := m.db.GetContext(ctx, &token, "SELECT resource_id, resource_type, expired_at FROM token WHERE resource_id = $1 AND resource_type = $2 ORDER BY id DESC", resourceId, "login")
		if checkTokenErr != nil {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "AccessToken not match with accessToken from db"
			helpers.Send(rw, helpers.ApiResponse(res))

			defer logrus.Errorf("Error Logs: %v", checkTokenErr)
			return
		}

		jakartaTimeZone, _ := time.LoadLocation("Asia/Bangkok")
		timeFormat := "2006-01-02 15:04:05"
		timeNow := time.Now().In(jakartaTimeZone).Format(timeFormat)

		if token.ExpiredAt.Format(timeFormat) < timeNow {
			res.StatCode = http.StatusBadRequest
			res.StatMsg = "AccessToken expired"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		cacheData := make(map[string]interface{})

		if metadataToken["role"] == "customer" {
			cacheData["id"] = customers.Id
			cacheData["email"] = customers.Email
			cacheData["role"] = customers.Role.Name
			cacheData["categorie"] = customers.Categorie.Name
		} else {
			cacheData["id"] = users.Id
			cacheData["email"] = users.Email
			cacheData["role"] = users.Role.Name
		}

		_, redisErr := packages.Redis(1).Hset(ctx, fmt.Sprintf("users:%d", users.Id), cacheData, time.Duration(helpers.ExpiredAt(1, "days")))
		if redisErr != nil {
			defer logrus.Errorf("Error Logs: %v", redisErr)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
