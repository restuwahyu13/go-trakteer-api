package models

import "time"

type DonationPricing struct {
	Id             int            `json:"id,omitempty" db:"id"`
	UserId         uint           `json:"user_id,omitempty" db:"user_id"`
	DonationTypeId uint           `json:"donation_type_id,omitempty" db:"donation_type_id"`
	Amount         uint64         `json:"amount,omitempty" db:"amount"`
	CreatedAt      time.Time      `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty" db:"updated_at"`
	User           *Users         `json:"user,omitempty"`
	DonationTypes  *DonationTypes `json:"donation_types,omitempty"`
}
