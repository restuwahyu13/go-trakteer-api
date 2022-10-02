package models

import "time"

type Balances struct {
	Id        uint      `db:"id"`
	Amount    uint64    `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
