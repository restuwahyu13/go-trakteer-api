package models

import "time"

type Wallets struct {
	Id        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	NoRek     uint      `json:"no_rek,omitempty" db:"no_rek"`
	BankName  string    `json:"bank_name,omitempty" db:"bank_name"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
