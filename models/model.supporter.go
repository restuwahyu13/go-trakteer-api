package models

import "time"

type Supporter struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Message   string    `json:"message" db:"message"`
	IsPrivate bool      `json:"is_private" db:"is_private"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
