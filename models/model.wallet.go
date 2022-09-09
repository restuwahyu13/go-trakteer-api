package models

import "time"

type Wallet struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	NoRek     uint64    `db:"no_rek"`
	BankName  string    `db:"bank_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
