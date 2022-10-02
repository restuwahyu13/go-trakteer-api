package models

import "time"

type Balances struct {
	Id        int       `json:"id" db:"id"`
	Amount    uint64    `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
