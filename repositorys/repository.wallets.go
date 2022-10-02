package repositorys

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/models"
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
	customers := models.Customers{}
	wallets := models.Wallets{}
	balances := models.Balances{}
	usersPayment := models.UsersPayment{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkWalletErrChan := make(chan error)
	checkUserErrChan := make(chan error)
	checkUserRowsChan := make(chan *sql.Rows)

	go func(checkWalletErrCh chan error, checkUserErrCh chan error, checkUserRowsCh chan *sql.Rows) {
		checkWalletErr := r.db.GetContext(ctx, &wallets, "SELECT id FROM wallet WHERE no_rek = $1", body.NoRek)
		checkWalletErrCh <- checkWalletErr

		checkUserRows, checkUserErr := r.db.QueryContext(ctx, `SELECT
		customers.id as customer_id, roles.id as role_id, roles.name as role_name FROM customers
		INNER JOIN roles ON customers.role_id = roles.id
		WHERE customers.id = $1 AND customers.active = $2 AND roles.name = $3`, body.CustomerId, "true", "customer")

		checkUserErrCh <- checkUserErr
		checkUserRowsCh <- checkUserRows

	}(checkWalletErrChan, checkUserErrChan, checkUserRowsChan)

	if err := <-checkWalletErrChan; err == nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("No rek %d already taken with other user", body.NoRek)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if err := <-checkUserErrChan; err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Customer for this id %d not exist", body.CustomerId)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if err := carta.Map(<-checkUserRowsChan, &customers); err != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Relation between table error: %v", err)
		defer logrus.Errorf("Error Logs: %v", err)
		return res
	}

	if customers.Role.Name != "customer" {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "You must be assign customer"
		return res
	}

	formatTimeZone, _ := time.LoadLocation("Asia/Bangkok")

	wallets.Name = body.Name
	wallets.NoRek = body.NoRek
	wallets.BankName = body.BankName
	wallets.CreatedAt = time.Now().In(formatTimeZone)

	var (
		walletId  int
		balanceId int
	)

	createWalletRows, createWalletErr := r.db.NamedQueryContext(ctx, "INSERT INTO wallet (name, no_rek, bank_name, created_at) VALUES (:name, :no_rek, :bank_name, :created_at) RETURNING id", &wallets)
	if createWalletRows.Next() {
		createWalletRows.Scan(&walletId)
	}

	if createWalletErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created customer wallet failed"
		defer logrus.Errorf("Error Logs: %v", createWalletErr)
		return res
	}

	balances.Amount = 0
	balances.CreatedAt = time.Now().In(formatTimeZone)

	createBalanceRows, createBalanceErr := r.db.NamedQueryContext(ctx, "INSERT INTO balance (amount, created_at) VALUES (:amount, :created_at) RETURNING id", &balances)
	if createBalanceRows.Next() {
		createBalanceRows.Scan(&balanceId)
	}

	if createBalanceErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created default balance failed"
		defer logrus.Errorf("Error Logs: %v", createBalanceErr)
		return res
	}

	usersPayment.CustomerId = body.CustomerId
	usersPayment.BalanceId = uint(balanceId)
	usersPayment.WalletId = uint(walletId)

	_, createUserPaymentErr := r.db.NamedQueryContext(ctx, "INSERT INTO users_payment (customer_id, balance_id, wallet_id) VALUES (:customer_id, :balance_id, :wallet_id)", &usersPayment)
	if createUserPaymentErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = "Created customer wallet failed"
		defer logrus.Errorf("Error Logs: %v", createUserPaymentErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Created customer wallet success"
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
