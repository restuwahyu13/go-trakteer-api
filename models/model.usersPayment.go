package models

type UsersPayment struct {
	Id         uint `db:"id"`
	CustomerId uint `db:"customer_id"`
	BalanceId  uint `db:"balance_id"`
	WalletId   uint `db:"wallet_id"`
	Customer   *Customers
	Balance    *Balances
	Wallet     *Wallets
}
