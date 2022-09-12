package models

import "time"

type Categories struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt any       `db:"deleted_at"`
	// Users     []Users
}
