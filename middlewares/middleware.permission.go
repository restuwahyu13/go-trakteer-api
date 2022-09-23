package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/packages"
)

type permission struct {
	roles []string
}

type usersMetadata struct {
	Id        uint   `redis:"id"`
	Email     string `redis:"email"`
	Role      string `redis:"role"`
	Categorie string `redis:"categorie"`
}

func NewMiddlewarePermission(roles ...string) *permission {
	return &permission{roles: roles}
}

func (h *permission) Middleware(next http.Handler) http.Handler {
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

		decodeToken, _ := packages.ParseToken(bearerToken[1])
		metadataToken := make(map[string]interface{})

		encoded, _ := json.Marshal(&decodeToken)
		json.Unmarshal(encoded, &metadataToken)

		redisRes, redisErr := packages.Redis(1).HGetAll(r.Context(), fmt.Sprintf("users:%s", metadataToken["email"]))

		if redisErr != nil {
			res.StatCode = http.StatusForbidden
			res.StatMsg = "Get caching user metadata failed"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		user := usersMetadata{}
		redisRes.Scan(&user)

		mappingRole := make(map[string]int)
		for i, v := range h.roles {
			mappingRole[v] = i + 1
		}

		if _, ok := mappingRole[user.Role]; !ok {
			res.StatCode = http.StatusForbidden
			res.StatMsg = "You role not allowed, access denield"
			helpers.Send(rw, helpers.ApiResponse(res))
			return
		}

		next.ServeHTTP(rw, r)
	})
}
