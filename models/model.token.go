package models

import "time"

type Token struct {
	Id           int       `json:"id,omitempty" db:"id"`
	ResourceId   uint      `json:"resource_id,omitempty" db:"resource_id"`
	ResourceType string    `json:"resource_type,omitempty" db:"resource_type"`
	AccessToken  string    `json:"access_token,omitempty" db:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty" db:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at,omitempty" db:"expired_at"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
}
