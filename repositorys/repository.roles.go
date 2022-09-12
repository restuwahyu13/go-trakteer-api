package repositorys

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

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

/**
* @description CreateRepository
**/

func (ctx *RolesRepository) CreateRepository(payload dtos.DTORoles) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	roles.Name = payload.Name
	checkRoleName := ctx.db.Get(&roles, "SELECT name FROM roles WHERE name = $1", payload.Name)

	if checkRoleName == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Role name already exist"
		return res
	}

	_, createdRoleErr := ctx.db.NamedQuery("INSERT INTO roles (name) VALUES (:name)", &roles)

	if createdRoleErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created new role failed"
		res.SqlError = createdRoleErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new role success"
	return res
}

/**
* @description GetAllRepository
**/

func (ctx *RolesRepository) GetAllRepository(query dtos.DTORolePagination) helpers.APIResponse {
	roles := []models.Roles{}
	res := helpers.APIResponse{}

	getAllRolesChan := make(chan error)
	countChan := make(chan int)

	go func() {
		getAllRoles := ctx.db.Select(&roles, fmt.Sprintf("SELECT * FROM roles ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)
		getAllRolesChan <- getAllRoles

		count := 0
		countRoles := ctx.db.QueryRowx("SELECT COUNT(id) FROM roles")
		countRoles.Scan(&countRoles)
		countChan <- count
	}()

	if <-getAllRolesChan != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = "Roles data not exist"
		res.SqlError = <-getAllRolesChan
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Roles already to use"
	res.Data = roles
	res.Pagination = helpers.Stringify(helpers.Pagination(query, <-countChan))
	return res
}

/**
* @description GetByIdRepository
**/

func (ctx *RolesRepository) GetByIdRepository(params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	getRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if getRoleId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", params.Id)
		res.SqlError = getRoleId
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Role already to use"
	res.Data = roles
	return res
}

/**
* @description DeleteByIdRepository
**/

func (ctx *RolesRepository) DeleteByIdRepository(params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if checkRoleId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", params.Id)
		res.SqlError = checkRoleId
		return res
	}

	_, deletedRoleErr := ctx.db.NamedQuery("DELETE FROM roles WHERE id = :id", params.Id)

	if deletedRoleErr != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Deleted role for this id %d failed", params.Id)
		res.SqlError = deletedRoleErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Deleted role for this id %d success", roles.ID)
	res.Data = roles
	return res
}

/**
* @description UpdatedByIdRepository
**/

func (ctx *RolesRepository) UpdatedByIdRepository(body dtos.DTORoles, params dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&roles, "SELECT * FROM roles WHERE id = $1", params.Id)

	if checkRoleId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", params.Id)
		res.SqlError = checkRoleId
		return res
	}

	roles.ID = uint(params.Id)
	roles.Name = body.Name
	roles.UpdatedAt = time.Now()

	_, updatedRoleErr := ctx.db.NamedQuery("UPDATE roles SET name = :name WHERE id = :id", &roles)

	if updatedRoleErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Updated old role failed"
		res.SqlError = updatedRoleErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated old role success"
	return res
}
