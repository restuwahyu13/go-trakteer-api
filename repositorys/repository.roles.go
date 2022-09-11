package repositorys

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
)

type RolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *RolesRepository {
	return &RolesRepository{db: db}
}

func (ctx *RolesRepository) CreateRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from CreateRepository",
	}

	return res
}

func (ctx *RolesRepository) GetAllRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from GetAllRepository",
	}

	return res
}

func (ctx *RolesRepository) GetByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from DeleteByIdRepository",
	}

	return res
}

func (ctx *RolesRepository) DeleteByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from DeleteByIdRepository",
	}

	return res
}

func (ctx *RolesRepository) UpdatedByIdRepository(payload dtos.DTOLogin) helpers.APIResponse {
	res := helpers.APIResponse{
		StatCode: http.StatusOK,
		StatMsg:  "Respon from UpdatedByIdRepository",
	}

	return res
}
