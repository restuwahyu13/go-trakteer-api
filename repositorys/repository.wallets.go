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

	tx := r.db.MustBeginTx(ctx, nil)

	checkWalletErrChan := make(chan error)
	checkUserErrChan := make(chan error)
	checkUserRowsChan := make(chan *sql.Rows)

	go func(checkWalletErrCh chan error, checkUserErrCh chan error, checkUserRowsCh chan *sql.Rows) {
		checkWalletErr := tx.GetContext(ctx, &wallets, "SELECT id FROM wallet WHERE no_rek = $1 OR name = $2", body.NoRek, body.Name)
		checkWalletErrCh <- checkWalletErr

		checkUserRows, checkUserErr := tx.QueryContext(ctx, `SELECT
		customers.id as customer_id, roles.id as role_id, roles.name as role_name FROM customers
		INNER JOIN roles ON customers.role_id = roles.id
		WHERE customers.id = $1 AND customers.active = $2 AND roles.name = $3`, body.CustomerId, "true", "customer")

		checkUserErrCh <- checkUserErr
		checkUserRowsCh <- checkUserRows

	}(checkWalletErrChan, checkUserErrChan, checkUserRowsChan)

	if err := <-checkWalletErrChan; err == nil {
		defer tx.Rollback()
		defer logrus.Errorf("Error Logs: %v", err)

		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("No rek %d already taken with other user or User already created wallet", body.NoRek)
		return res
	}

	if err := <-checkUserErrChan; err != nil {
		defer tx.Rollback()
		defer logrus.Errorf("Error Logs: %v", err)

		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Customer for this id %d not exist", body.CustomerId)
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

	createWalletRows := tx.QueryRowContext(ctx, "INSERT INTO wallet (name, no_rek, bank_name, created_at) VALUES ($1, $2, $3, $4) RETURNING id", wallets.Name, wallets.NoRek, wallets.BankName, wallets.CreatedAt)
	if createWalletRows.Err() == nil {
		createWalletRows.Scan(&walletId)
	}

	if err := createWalletRows.Err(); err != nil {
		defer tx.Rollback()
		defer logrus.Errorf("Error Logs: %v", err)

		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created customer wallet failed"

		return res
	}

	balances.Amount = 0
	balances.CreatedAt = time.Now().In(formatTimeZone)

	createBalanceRows := tx.QueryRowContext(ctx, "INSERT INTO balance (amount, created_at) VALUES ($1, $2) RETURNING id", balances.Amount, balances.CreatedAt)
	if createBalanceRows.Err() == nil {
		createBalanceRows.Scan(&balanceId)
	}

	if err := createBalanceRows.Err(); err != nil {
		defer tx.Rollback()
		defer logrus.Errorf("Error Logs: %v", err)

		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created default balance failed"
		return res
	}

	usersPayment.CustomerId = body.CustomerId
	usersPayment.BalanceId = uint(balanceId)
	usersPayment.WalletId = uint(walletId)

	createUserPaymentRows := tx.QueryRowContext(ctx, "INSERT INTO users_payment (customer_id, balance_id, wallet_id) VALUES ($1, $2, $3)", usersPayment.CustomerId, usersPayment.BalanceId, usersPayment.WalletId)
	if err := createUserPaymentRows.Err(); err != nil {
		defer tx.Rollback()
		defer logrus.Errorf("Error Logs: %v", err)

		res.StatCode = http.StatusForbidden
		res.StatMsg = "Created customer wallet failed"
		return res
	}

	if err := tx.Commit(); err != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Commit transaction failed"
		defer logrus.Errorf("Error Logs: %v", err)
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
	wallets := models.Wallets{}

	ctx, cancel := context.WithTimeout(ctx, min)
	defer cancel()

	checkWalletIdErr := r.db.GetContext(ctx, &wallets, "SELECT * FROM wallet WHERE id = $1", params.Id)

	if checkWalletIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Wallet data for this id %d not exist", params.Id)
		defer logrus.Errorf("Error Logs: %v", checkWalletIdErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Wallet already to use"
	res.Data = wallets
	return res
}

/**
* @description UpdateWalletsByIdRepository
**/

func (r *walletsRepository) UpdateByIdRepository(ctx context.Context, body *dtos.DTOWalletsUpdate, params *dtos.DTOWalletsById) helpers.APIResponse {
	res := helpers.APIResponse{}
	wallets := models.Wallets{}

	checkWalletIdErr := r.db.GetContext(ctx, &wallets, "SELECT id FROM wallet WHERE id = $1", params.Id)

	if checkWalletIdErr != nil {
		res.StatCode = http.StatusBadRequest
		res.StatMsg = fmt.Sprintf("Wallet data for this id %d not exist", params.Id)
		defer logrus.Errorf("Error Logs: %v", checkWalletIdErr)
		return res
	}

	wallets.Id = int(params.Id)
	wallets.Name = body.Name
	wallets.NoRek = body.NoRek
	wallets.BankName = body.BankName
	wallets.UpdatedAt = time.Now().Local()

	_, updateWalletErr := r.db.NamedQuery("UPDATE wallet SET name = :name, no_rek = :no_rek, bank_name = :bank_name WHERE id = :id", &wallets)

	if updateWalletErr != nil {
		res.StatCode = http.StatusForbidden
		res.StatMsg = "Updated wallet failed"
		defer logrus.Errorf("Error Logs: %v", checkWalletIdErr)
		return res
	}

	res.StatCode = http.StatusOK
	res.StatMsg = "Updated wallet success"
	return res
}
