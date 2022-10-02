package models

import "time"

type Withdraw struct {
	Id        int       `json:"id,omitempty" db:"id"`
	BalanceId uint      `json:"balance_id,omitempty" db:"balance_id"`
	Amount    uint64    `json:"amount,omitempty" db:"amount"`
	DateTime  time.Time `json:"date_time,omitempty" db:"date_time"`
	Fee       uint      `json:"fee,omitempty" db:"fee"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	Balance   *Balances `json:"balance,omitempty"`
}
