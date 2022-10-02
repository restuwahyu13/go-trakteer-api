package models

import (
	"time"
)

type Streams struct {
	Id            int       `json:"id,omitempty" db:"id"`
	Notification  string    `json:"notification,omitempty" db:"notification"`
	Leaderboard   string    `json:"leaderboard,omitempty" db:"leaderboard"`
	LastSupporter string    `json:"last_supporter,omitempty" db:"last_supporter"`
	Target        string    `json:"target,omitempty" db:"target"`
	RunningText   string    `json:"running_text,omitempty" db:"running_text"`
	Subathon      string    `json:"subathon,omitempty" db:"subathon"`
	Qrcode        string    `json:"qrcode,omitempty" db:"qrcode"`
	Key           string    `json:"key,omitempty" db:"key"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Goals         *[]Goal   `json:"goals,omitempty"`
}
