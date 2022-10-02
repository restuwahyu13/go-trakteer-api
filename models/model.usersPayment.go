package models

type relationCustomer struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type relationBalance struct {
	Amount uint64 `json:"amount"`
}

type relationWallet struct {
	Name     string `json:"name"`
	NoRek    uint   `json:"no_rek"`
	BankName string `json:"bank_name"`
}

type UsersPayment struct {
	Id         int               `json:"id" db:"id"`
	CustomerId uint              `json:"customer_id" db:"customer_id"`
	BalanceId  uint              `json:"balance_id" db:"balance_id"`
	WalletId   uint              `json:"wallet_id" db:"wallet_id"`
	Customer   *relationCustomer `json:"customer,omitempty"`
	Balance    *relationBalance  `json:"balance,omitempty"`
	Wallet     *relationWallet   `json:"wallet,omitempty"`
}
