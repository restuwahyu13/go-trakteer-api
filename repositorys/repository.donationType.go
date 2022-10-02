package repositorys

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
)

type DonationTypeRepository = interfaces.IDonationTypeRepository
type donationTypeRepository struct {
	db *sqlx.DB
}

func NewDonationTypeRepository(db *sqlx.DB) *donationTypeRepository {
	return &donationTypeRepository{db: db}
}

/*
@description CreateRepository
*/

func (r *donationTypeRepository) CreateRepository(ctx context.Context, body *dtos.DTODonationType) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "CreateRepository"
	res.Data = body
	return res
}

/*
@description GetAllRepository
*/

func (r *donationTypeRepository) GetAllRepository(ctx context.Context, query *dtos.DTODonationTypePagination) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "GetAllRepository"
	return res
}

/*
@description GetByIdRepository
*/

func (r *donationTypeRepository) GetByIdRepository(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "GetByIdRepository"
	return res
}

/*
@description DeleteByIdRepository
*/

func (r *donationTypeRepository) DeleteByIdRepository(ctx context.Context, params *dtos.DTODonationTypeId) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "DeleteByIdRepository"
	return res
}

/*
@description UpdateByIdRepository
*/

func (r *donationTypeRepository) UpdateByIdRepository(ctx context.Context, body *dtos.DTODonationType, params *dtos.DTODonationTypeId) helpers.APIResponse {
	res := helpers.APIResponse{}

	res.StatCode = http.StatusOK
	res.StatMsg = "UpdateByIdRepository"
	return res
}
