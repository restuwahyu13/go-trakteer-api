package models

import "time"

type Goal struct {
	Id            int       `json:"id" db:"id"`
	StreamId      uint      `json:"stream_id" db:"stream_id"`
	Name          string    `json:"name" db:"name"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"start_date"`
	TargetBalance uint64    `json:"target_balance" db:"target_balance"`
	Progress      string    `json:"progress" db:"progress"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Payments      []Payments
}
