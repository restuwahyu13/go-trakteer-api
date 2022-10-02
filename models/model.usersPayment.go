package models

type UsersPayment struct {
	Id         int  `json:"id" db:"id"`
	CustomerId uint `json:"customer_id" db:"customer_id"`
	BalanceId  uint `json:"balance_id" db:"balance_id"`
	WalletId   uint `json:"wallet_id" db:"wallet_id"`
	Customer   *Customers
	Balance    *Balances
	Wallet     *Wallets
}
