package dtos

import "time"

type DTOWalletsCreate struct {
	Name       string    `validate:"required" json:"name"`
	NoRek      uint      `validate:"required,numeric" json:"no_rek"`
	BankName   string    `validate:"required" json:"bank_name"`
	CustomerId uint      `validate:"required,numeric" json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DTOWalletsById struct {
	Id uint `validate:"required" json:"id"`
}

type DTOWalletsUpdate struct {
	Name      string    `validate:"required" json:"name"`
	NoRek     uint      `validate:"required" json:"no_rek"`
	BankName  string    `validate:"required" json:"bank_name"`
	Amount    uint64    `validate:"required" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
