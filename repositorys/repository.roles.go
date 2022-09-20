package repositorys

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/models"
)

type RolesRepository = interfaces.IRolesRepository
type rolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *rolesRepository {
	return &rolesRepository{db: db}
}

/**
* @description CreateRepository
**/

func (r *rolesRepository) CreateRepository(ctx context.Context, body *dtos.DTORoles) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	roles.Name = body.Name
	checkRoleNameErr := r.db.GetContext(ctx, &roles, "SELECT name FROM roles WHERE name = $1", body.Name)

	if checkRoleNameErr == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Role name already exist"
		defer logrus.Errorf("Error Logs: %v", checkRoleNameErr)
		return res
	}

	_, createdRoleErr := r.db.NamedQueryContext(ctx, "INSERT INTO roles (name) VALUES (:name)", &roles)

	if createdRoleErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created new role failed"
		defer logrus.Errorf("Error Logs: %v", createdRoleErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new role success"
	return res
}

/**
* @description GetAllRepository
**/

func (r *rolesRepository) GetAllRepository(ctx context.Context, query *dtos.DTORolePagination) helpers.APIResponse {
	roles := []models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, max)
	defer cancel()

	getAllRolesChan := make(chan error)
	countChan := make(chan int)

	go func() {
		getAllRoles := r.db.SelectContext(ctx, &roles, fmt.Sprintf("SELECT * FROM roles ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)
		getAllRolesChan <- getAllRoles

		count := 0
		countRoles := r.db.QueryRowContext(ctx, "SELECT COUNT(id) FROM roles")
		countRoles.Scan(&count)
		countChan <- count
	}()

	if err := <-getAllRolesChan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Roles data not exist"
		defer logrus.Errorf("Error Logs: %v", err)
		defer close(getAllRolesChan)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Roles already to use"
	res.Data = roles
	res.Pagination = helpers.Stringify(helpers.Pagination(query, <-countChan))
	defer close(countChan)
	return res
}

/**
* @description GetByIdRepository
**/

func (r *rolesRepository) GetByIdRepository(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	roles.Id = params.Id
	getRoleIdErr := r.db.GetContext(ctx, &roles, "SELECT * FROM roles WHERE id = $1", roles.Id)

	if getRoleIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d not exist", roles.Id)
		defer logrus.Errorf("Error Logs: %v", getRoleIdErr)
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

func (r *rolesRepository) DeleteByIdRepository(ctx context.Context, params *dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	roles.Id = params.Id
	checkRoleIdErr := r.db.GetContext(ctx, &roles, "SELECT id FROM roles WHERE id = $1", roles.Id)

	if checkRoleIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", roles.Id)
		defer logrus.Errorf("Error Logs: %v", checkRoleIdErr)
		return res
	}

	deletedTime := time.Now().Local()
	roles.DeletedAt = &deletedTime

	_, deletedRoleErr := r.db.NamedQueryContext(ctx, "UPDATE roles SET deleted_at = :deleted_at WHERE id = :id", &roles)

	if deletedRoleErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Deleted role for this id %d failed", roles.Id)
		defer logrus.Errorf("Error Logs: %v", deletedRoleErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Deleted role for this id %d success", roles.Id)
	res.Data = roles
	return res
}

/**
* @description UpdatedByIdRepository
**/

func (r *rolesRepository) UpdatedByIdRepository(ctx context.Context, body *dtos.DTORoles, params *dtos.DTORolesById) helpers.APIResponse {
	roles := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	roles.Id = params.Id
	checkRoleIdErr := r.db.GetContext(ctx, &roles, "SELECT id FROM roles WHERE id = $1", roles.Id)

	if checkRoleIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", roles.Id)
		defer logrus.Errorf("Error Logs: %v", checkRoleIdErr)
		return res
	}

	roles.Name = body.Name
	roles.UpdatedAt = time.Now()

	_, updatedRoleErr := r.db.NamedQueryContext(ctx, "UPDATE roles SET name = :name WHERE id = :id", &roles)

	if updatedRoleErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Updated role data for this id %v success", roles.Id)
		defer logrus.Errorf("Error Logs: %v", updatedRoleErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Updated role data for this id %v success", roles.Id)
	return res
}
