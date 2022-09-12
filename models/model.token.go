package models

import "time"

type Token struct {
	ID           uint      `db:"id"`
	ResourceId   uint      `db:"resource_id"`
	ResourceType string    `db:"resource_type"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	ExpiredAt    time.Time `db:"expired_at"`
	CreatedAt    time.Time `db:"created_at"`
}
