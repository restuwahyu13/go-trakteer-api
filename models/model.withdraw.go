package models

import "time"

type Withdraw struct {
	Id        int       `json:"id" db:"id"`
	BalanceId uint      `json:"balance_id" db:"balance_id"`
	Amount    uint64    `json:"amount" db:"amount"`
	DateTime  time.Time `json:"date_time" db:"date_time"`
	Fee       uint      `json:"fee" db:"fee"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Balance   *Balances
}
