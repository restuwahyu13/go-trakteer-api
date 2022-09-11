package repositorys

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
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

func (ctx *CategoriesRepository) GetAllRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from GetAllRepository",
	}

	return res
}

func (ctx *CategoriesRepository) GetByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
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
