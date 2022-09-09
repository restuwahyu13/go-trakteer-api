package models

import "time"

type Supporter struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	Message   string    `db:"message"`
	IsPrivate bool      `db:"is_private"`
	CreatedAt time.Time `db:"created_at"`
}
