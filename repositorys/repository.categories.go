package repositorys

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/models"
)

type CategoriesRepository struct {
	db *sqlx.DB
}

func NewCategoriesRepository(db *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

/**
* @description CreateRepository
**/

func (ctx *CategoriesRepository) CreateRepository(body dtos.DTOCategories) helpers.APIResponse {
	categories := models.Roles{}
	res := helpers.APIResponse{}

	categories.Name = body.Name
	checkRoleName := ctx.db.Get(&categories, "SELECT name FROM categories WHERE name = $1", body.Name)

	if checkRoleName == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Categorie name already exist"
		return res
	}

	_, err := ctx.db.NamedQuery("INSERT INTO categories (name) VALUES (:name)", &categories)

	if err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created new categorie failed"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new categorie success"
	return res
}

/**
* @description GetAllRepository
**/

func (ctx *CategoriesRepository) GetAllRepository(query dtos.DTOCategoriesPagination) helpers.APIResponse {
	categories := []models.Categories{}
	res := helpers.APIResponse{}

	getAllCategories := ctx.db.Select(&categories, helpers.Strings("SELECT * FROM categories ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)

	if getAllCategories != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = "Categories data not exist"
		return res
	}

	count := 0
	countCategories := ctx.db.QueryRowx("SELECT COUNT(id) FROM categories")
	countCategories.Scan(&countCategories)

	res.StatCode = http.StatusOK
	res.StatMsg = "Categories already to use"
	res.Data = categories
	res.Pagination = helpers.Stringify(helpers.Pagination(query, count))
	return res
}

/**
* @description GetByIdRepository
**/

func (ctx *CategoriesRepository) GetByIdRepository(params dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Categories{}
	res := helpers.APIResponse{}

	getRoleId := ctx.db.Get(&catagories, "SELECT * FROM catagories WHERE id = $1", params.Id)

	if getRoleId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Categorie data for this id %d, not exist", params.Id)
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

func (ctx *CategoriesRepository) DeleteByIdRepository(params dtos.DTOCategoriesId) helpers.APIResponse {
	categories := models.Categories{}
	res := helpers.APIResponse{}

	checkCategorieId := ctx.db.Get(&categories, "SELECT * FROM categories WHERE id = $1", params.Id)

	if checkCategorieId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Role data for this id %d, not exist", params.Id)
		return res
	}

	_, err := ctx.db.NamedQuery("DELETE FROM categories WHERE id = :id", params.Id)

	if err != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Deleted categorie for this id %s failed", params.Id)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = helpers.Strings("Deleted categorie for this id %s success", categories.ID)
	res.Data = categories
	return res
}

/**
* @description UpdatedByIdRepository
**/

func (ctx *CategoriesRepository) UpdatedByIdRepository(body dtos.DTOCategories, params dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Roles{}
	res := helpers.APIResponse{}

	checkRoleId := ctx.db.Get(&catagories, "SELECT * FROM catagories WHERE id = $1", params.Id)

	if checkRoleId != nil {
		res.StatCode = http.StatusNotFound
		res.StatMsg = helpers.Strings("Role data for this id %d, not exist", params.Id)
		return res
	}

	catagories.ID = uint(params.Id)
	catagories.Name = body.Name
	catagories.UpdatedAt = time.Now()

	_, err := ctx.db.NamedQuery("UPDATE catagories SET name = :name WHERE id = :id", &catagories)

	if err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Updated old categorie failed"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated old categorie success"
	return res
}
