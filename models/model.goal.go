package models

import "time"

type Goal struct {
	ID            uint      `db:"id"`
	StreamId      uint      `db:"stream_id"`
	Name          string    `db:"name"`
	StartDate     time.Time `db:"start_date"`
	EndDate       time.Time `db:"start_date"`
	TargetBalance uint64    `db:"target_balance"`
	Progress      string    `db:"progress"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Stream        Streams
}
