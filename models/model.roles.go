package models

import "time"

type Roles struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"created_at"`
	Users     []Users
}
