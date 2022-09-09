package models

import "time"

type Token struct {
	ID           uint      `db:"id"`
	ResourceID   uint      `db:"resource_id"`
	ResourceBy   uint      `db:"resource_by"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	ExpiredAt    time.Time `db:"expired_at"`
	CreatedAt    time.Time `db:"created_at"`
}
