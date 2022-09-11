package models

import "time"

type Withdraw struct {
	Id        uint      `db:"id"`
	BalanceID uint      `db:"balance_id"`
	Amount    uint64    `db:"amount"`
	DateTime  time.Time `db:"date_time"`
	Fee       uint      `db:"fee"`
	CreatedAt time.Time `db:"created_at"`
	Balance   Balance
}
