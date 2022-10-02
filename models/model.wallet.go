package models

import "time"

type Wallets struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	NoRek     uint      `json:"no_rek" db:"no_rek"`
	BankName  string    `json:"bank_name" db:"bank_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
