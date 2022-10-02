package models

import "time"

type DonationPricing struct {
	Id             int       `json:"id" db:"id"`
	UserId         uint      `json:"user_id" db:"user_id"`
	DonationTypeId uint      `json:"donation_type_id" db:"donation_type_id"`
	Amount         uint64    `json:"amount" db:"amount"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	User           Users
	DonationTypes  DonationTypes
}
