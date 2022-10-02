package models

import "time"

type Supporter struct {
	Id        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	Message   string    `json:"message,omitempty" db:"message"`
	IsPrivate bool      `json:"is_private,omitempty" db:"is_private"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
