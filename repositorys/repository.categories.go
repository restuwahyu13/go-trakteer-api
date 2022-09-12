package repositorys

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

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

func (ctx *CategoriesRepository) CreateRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from CreateRepository",
	}

	return res
}

func (ctx *CategoriesRepository) GetAllRepository(payload dtos.DTOCategoriesPagination) helpers.APIResponse {
	categories := []models.Categories{}
	res := helpers.APIResponse{}

	row := ctx.db.Select(&categories, helpers.Strings("SELECT * FROM categories ORDER BY id %s LIMIT $1 OFFSET $2", payload.Sort), payload.Limit, payload.Offset)
	if row != nil {
		defer logrus.Error("Query Error: %v", row.Error())

		res.StatCode = http.StatusNotFound
		res.StatMsg = "Categories data not exist"
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Categories already to use"
	res.Data = categories
	res.Pagination = helpers.Stringify(helpers.Pagination(payload, len(categories)))
	return res
}

func (ctx *CategoriesRepository) GetByIdRepository(payload dtos.DTORolesById) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from DeleteByIdRepository",
	}

	return res
}

func (ctx *CategoriesRepository) DeleteByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from DeleteByIdRepository",
	}

	return res
}

func (ctx *CategoriesRepository) UpdatedByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from UpdatedByIdRepository",
	}

	return res
}
