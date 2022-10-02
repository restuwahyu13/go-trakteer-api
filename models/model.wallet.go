package models

import "time"

type Wallets struct {
	Id        uint      `db:"id"`
	Name      string    `db:"name"`
	NoRek     uint      `db:"no_rek"`
	BankName  string    `db:"bank_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
