package models

import "time"

type DonationPricing struct {
	Id             uint      `db:"id"`
	UserId         uint      `db:"user_id"`
	DonationTypeId uint      `db:"donation_type_id"`
	Amount         uint64    `db:"amount"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	User           Users
	DonationTypes  DonationTypes
}
