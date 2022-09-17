package repositorys

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/models"
)

type CategoriesRepository = interfaces.ICategoriesRepository
type categoriesRepository struct {
	db *sqlx.DB
}

func NewCategoriesRepository(db *sqlx.DB) *categoriesRepository {
	return &categoriesRepository{db: db}
}

/**
* @description CreateRepository
**/

func (ctx *categoriesRepository) CreateRepository(body *dtos.DTOCategories) helpers.APIResponse {
	categories := models.Roles{}
	res := helpers.APIResponse{}

	categories.Name = body.Name
	checkRoleName := ctx.db.Get(&categories, "SELECT name FROM categories WHERE name = $1", body.Name)

	if checkRoleName == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Categorie name already exist"
		return res
	}

	_, createdCategorieErr := ctx.db.NamedQuery("INSERT INTO categories (name) VALUES (:name)", &categories)

	if createdCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created new categorie failed"
		res.QueryError = createdCategorieErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new categorie success"
	return res
}

/**
* @description GetAllRepository
**/

func (ctx *categoriesRepository) GetAllRepository(query *dtos.DTOCategoriesPagination) helpers.APIResponse {
	categories := []models.Categories{}
	res := helpers.APIResponse{}

	getAllCategoriesChan := make(chan error)
	countChan := make(chan int)

	go (func() {
		getAllCategories := ctx.db.Select(&categories, fmt.Sprintf("SELECT * FROM categories ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)
		getAllCategoriesChan <- getAllCategories

		count := 0
		countCategories := ctx.db.QueryRowx("SELECT COUNT(id) FROM categories")
		countCategories.Scan(&count)
		countChan <- count
	})()

	if <-getAllCategoriesChan != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Categories data not exist"
		res.QueryError = <-getAllCategoriesChan
		defer close(getAllCategoriesChan)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Categories already to use"
	res.Data = categories
	res.Pagination = helpers.Stringify(helpers.Pagination(query, <-countChan))
	defer close(countChan)
	return res
}

/**
* @description GetByIdRepository
**/

func (ctx *categoriesRepository) GetByIdRepository(params *dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Categories{}
	res := helpers.APIResponse{}

	getRoleId := ctx.db.Get(&catagories, "SELECT * FROM catagories WHERE id = $1", params.Id)

	if getRoleId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Categorie data for this id %d, not exist", params.Id)
		res.QueryError = getRoleId
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Categorie already to use"
	res.Data = catagories
	return res
}

/**
* @description DeleteByIdRepository
**/

func (ctx *categoriesRepository) DeleteByIdRepository(params *dtos.DTOCategoriesId) helpers.APIResponse {
	categories := models.Categories{}
	res := helpers.APIResponse{}

	checkCategorieId := ctx.db.Get(&categories, "SELECT * FROM categories WHERE id = $1", params.Id)

	if checkCategorieId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", params.Id)
		res.QueryError = checkCategorieId
		return res
	}

	_, deletedCategorieErr := ctx.db.NamedQuery("DELETE FROM categories WHERE id = :id", params.Id)

	if deletedCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Deleted categorie for this id %d failed", params.Id)
		res.QueryError = deletedCategorieErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Deleted categorie for this id %d success", categories.Id)
	res.Data = categories
	return res
}

/**
* @description UpdatedByIdRepository
**/

func (ctx *categoriesRepository) UpdatedByIdRepository(body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&catagories, "SELECT * FROM catagories WHERE id = $1", params.Id)

	if checkRoleId != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Role data for this id %d, not exist", params.Id)
		res.QueryError = checkRoleId
		return res
	}

	catagories.ID = uint(params.Id)
	catagories.Name = body.Name
	catagories.UpdatedAt = time.Now()

	_, updatedCategorieErr := ctx.db.NamedQuery("UPDATE catagories SET name = :name WHERE id = :id", &catagories)

	if updatedCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Updated old categorie failed"
		res.QueryError = updatedCategorieErr
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated old categorie success"
	return res
}
