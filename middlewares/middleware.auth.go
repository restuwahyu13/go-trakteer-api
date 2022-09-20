package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

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
		token := models.Token{}

		decodeToken, _ := packages.ParseToken(bearerToken[1])
		metadataToken := make(map[string]interface{})

		encoded, _ := json.Marshal(&decodeToken)
		json.Unmarshal(encoded, &metadataToken)

		checkUserEmailErr := m.db.GetContext(ctx, &users, "SELECT id FROM users WHERE email = $1", metadataToken["email"])
		if checkUserEmailErr != nil {
			res.StatCode = http.StatusUnauthorized
			res.StatMsg = "Metadata accessToken not match with metadataToken from db"
			helpers.Send(rw, helpers.ApiResponse(res))

			defer logrus.Errorf("Error Logs: %v", checkUserEmailErr)
			return
		}

		checkTokenErr := m.db.GetContext(ctx, &token, "SELECT resource_id, resource_type, expired_at FROM token WHERE resource_id = $1 AND resource_type = $2 ORDER BY id DESC", users.Id, "login")
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
		}

		next.ServeHTTP(rw, r)
	})
}
