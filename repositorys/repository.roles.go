package repositorys

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/models"
)

type RolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *RolesRepository {
	return &RolesRepository{db: db}
}

func (ctx *RolesRepository) CreateRepository(payload dtos.DTORoles) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	roles.Name = payload.Name
	checkRoleName := ctx.db.Get(&roles, "SELECT name FROM roles WHERE name = $1", payload.Name)

	if checkRoleName == nil {
		defer logrus.Error("Query Error: %v", checkRoleName.Error())

		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Role name already exist"
		return res
	}

	_, err := ctx.db.NamedQuery("INSERT INTO roles (name) VALUES (:name)", &roles)

	if err != nil {
		defer logrus.Error("Query Error: %v", err)

		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created new role failed"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new role success"
	return res
}

func (ctx *RolesRepository) GetAllRepository(query dtos.DTORolePagination) helpers.APIResponse {
	roles := []models.Roles{}
	res := helpers.APIResponse{}

	row := ctx.db.Select(&roles, helpers.Strings("SELECT * FROM roles ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)

	if row != nil {
		defer logrus.Error("Query Error: %v", row.Error())

		res.StatCode = http.StatusNotFound
		res.StatMsg = "Roles data not exist"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Roles already to use"
	res.Data = roles
	res.Pagination = helpers.Stringify(helpers.Pagination(query, len(roles)))
	return res
}

func (ctx *RolesRepository) GetByIdRepository(params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	getRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if getRoleId != nil {
		defer logrus.Error("Query Error: %v", getRoleId.Error())

		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Role data for this id %d, not exist", params.Id)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Role already to use"
	res.Data = roles
	return res
}

func (ctx *RolesRepository) DeleteByIdRepository(params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if checkRoleId != nil {
		defer logrus.Error("Query Error: %v", checkRoleId.Error())

		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Role data for this id %d, not exist", params.Id)
		return res
	}

	_, err := ctx.db.NamedQuery("DELETE FROM roles WHERE id = :id", params.Id)

	if err != nil {
		defer logrus.Error("Query Error: %v", err)

		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Deleted role for this id %s failed", params.Id)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = helpers.Strings("Deleted role for this id %s success", roles.ID)
	res.Data = roles
	return res
}

func (ctx *RolesRepository) UpdatedByIdRepository(body dtos.DTORoles, params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if checkRoleId != nil {
		defer logrus.Error("Query Error: %v", checkRoleId.Error())

		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Role data for this id %d, not exist", params.Id)
		return res
	}

	roles.ID = uint(params.Id)
	roles.Name = body.Name
	roles.UpdatedAt = time.Now()

	_, err := ctx.db.NamedQuery("UPDATE roles SET name = :name WHERE id = :id", &roles)

	if err != nil {
		defer logrus.Error("Query Error: %v", err)

		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Updated old role failed"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated old role success"
	return res
}
