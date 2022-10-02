package models

import "time"

type Goal struct {
	Id            int       `json:"id,omitempty" db:"id"`
	StreamId      uint      `json:"stream_id,omitempty" db:"stream_id"`
	Name          string    `json:"name,omitempty" db:"name"`
	StartDate     time.Time `json:"start_dat,omitempty" db:"start_date"`
	EndDate       time.Time `json:"end_date,omitempty" db:"start_date"`
	TargetBalance uint64    `json:"target_balance" db:"target_balance"`
	Progress      int       `json:"progress" db:"progress"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Payments      []Payments
}
