package models

import "time"

type Roles struct {
	ID        uint      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt any       `json:"deleted_at" db:"deleted_at"`
	// Users     []Users
}
