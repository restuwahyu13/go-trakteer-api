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

func (r *categoriesRepository) CreateRepository(ctx context.Context, body *dtos.DTOCategories) helpers.APIResponse {
	categories := models.Roles{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	categories.Name = body.Name
	checkRoleNameErr := r.db.GetContext(ctx, &categories, "SELECT name FROM categories WHERE name = $1", body.Name)

	if checkRoleNameErr == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Categorie name already exist"
		defer logrus.Errorf("Error Logs: %v", checkRoleNameErr)
		return res
	}

	_, createdCategorieErr := r.db.NamedQueryContext(ctx, "INSERT INTO categories (name) VALUES (:name)", &categories)

	if createdCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created new categorie failed"
		defer logrus.Errorf("Error Logs: %v", createdCategorieErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created new categorie success"
	return res
}

/**
* @description GetAllRepository
**/

func (r *categoriesRepository) GetAllRepository(ctx context.Context, query *dtos.DTOCategoriesPagination) helpers.APIResponse {
	categories := []models.Categories{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, max)
	defer cancel()

	getAllCategoriesChan := make(chan error)
	countChan := make(chan int)

	go func(getAllCategoriesCh chan error, countCh chan int) {
		getAllCategories := r.db.SelectContext(ctx, &categories, fmt.Sprintf("SELECT * FROM categories ORDER BY id %s LIMIT $1 OFFSET $2", query.Sort), query.Limit, query.Offset)
		getAllCategoriesCh <- getAllCategories

		count := 0
		countCategories := r.db.QueryRowContext(ctx, "SELECT COUNT(id) FROM categories")
		countCategories.Scan(&count)
		countCh <- count
	}(getAllCategoriesChan, countChan)

	if err := <-getAllCategoriesChan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Categories data not exist"
		defer logrus.Errorf("Error Logs: %v", err)
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

func (r *categoriesRepository) GetByIdRepository(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Categories{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	catagories.Id = params.Id
	getRoleIdErr := r.db.GetContext(ctx, &catagories, "SELECT * FROM categories WHERE id = $1", catagories.Id)

	if getRoleIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("CategorieId for this id %d not exist", catagories.Id)
		defer logrus.Errorf("Error Logs: %v", getRoleIdErr)
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

func (r *categoriesRepository) DeleteByIdRepository(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse {
	categories := models.Categories{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	categories.Id = params.Id
	checkCategorieIdErr := r.db.GetContext(ctx, &categories, "SELECT * FROM categories WHERE id = $1", categories.Id)

	if checkCategorieIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Categorie data for this id %d, not exist", categories.Id)
		defer logrus.Errorf("Error Logs: %v", checkCategorieIdErr)
		return res
	}

	deletedTime := time.Now().Local()
	categories.DeletedAt = &deletedTime

	_, deletedCategorieErr := r.db.NamedQueryContext(ctx, "UPDATE categories SET deleted_at = :deleted_at WHERE id = :id", &categories)

	if deletedCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Deleted categorie for this id %d failed", categories.Id)
		defer logrus.Errorf("Error Logs: %v", deletedCategorieErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Deleted categorie for this id %d success", categories.Id)
	return res
}

/**
* @description UpdatedByIdRepository
**/

func (r *categoriesRepository) UpdatedByIdRepository(ctx context.Context, body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse {
	catagories := models.Categories{}
	res := helpers.APIResponse{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	catagories.Id = params.Id
	checkCategorieIdErr := r.db.GetContext(ctx, &catagories, "SELECT id FROM categories WHERE id = $1", catagories.Id)

	if checkCategorieIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Categorie data for this id %d not exist", catagories.Id)
		defer logrus.Errorf("Error Logs: %v", checkCategorieIdErr)
		return res
	}

	catagories.Name = body.Name
	catagories.UpdatedAt = time.Now().Local()

	_, updatedCategorieErr := r.db.NamedQueryContext(ctx, "UPDATE categories SET name = :name WHERE id = :id", &catagories)

	if updatedCategorieErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = fmt.Sprintf("Updated categorie data for this id %v failed", catagories.Id)
		defer logrus.Errorf("Error Logs: %v", updatedCategorieErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = fmt.Sprintf("Updated categorie data for this id %v success", catagories.Id)
	return res
}
