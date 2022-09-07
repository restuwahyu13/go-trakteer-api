package models

import "time"

type Balance struct {
	Id        int       `db:"id"`
	Amount    uint64    `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
