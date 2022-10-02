package models

import "time"

type Token struct {
	Id           int       `json:"id" db:"id"`
	ResourceId   uint      `json:"resource_id" db:"resource_id"`
	ResourceType string    `json:"resource_type" db:"resource_type"`
	AccessToken  string    `json:"access_token" db:"access_token"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at" db:"expired_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
