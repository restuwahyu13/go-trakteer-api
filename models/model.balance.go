package models

import "time"

type Balances struct {
	Id        int       `json:"id,omitempty" db:"id"`
	Amount    uint64    `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
