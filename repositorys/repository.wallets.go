package repositorys

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
)

type WalletsRepository = interfaces.IWalletsRepository
type walletsRepository struct {
	db *sqlx.DB
}

func NewWalletsRepository(db *sqlx.DB) *walletsRepository {
	return &walletsRepository{db: db}
}

/**
* @description CreateWalletsRepository
**/

func (r *walletsRepository) CreateRepository(ctx context.Context, body *dtos.DTOWalletsCreate) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "CreateWalletsRepository"
	return res
}

/**
* @description GetWalletsById
**/

func (r *walletsRepository) GetByIdRepository(ctx context.Context, params *dtos.DTOWalletsById) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "GetWalletsById"
	return res
}

/**
* @description UpdateWalletsByIdRepository
**/

func (r *walletsRepository) UpdateByIdRepository(ctx context.Context, body *dtos.DTOWalletsUpdate, params *dtos.DTOWalletsById) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "UpdateWalletsByIdRepository"
	return res
}
