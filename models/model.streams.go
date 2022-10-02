package models

import (
	"encoding/json"
	"time"
)

type Streams struct {
	Id            int             `json:"id" db:"id"`
	Notification  json.RawMessage `json:"notification" db:"notification"`
	Leaderboard   json.RawMessage `json:"leaderboard" db:"leaderboard"`
	LastSupporter json.RawMessage `json:"last_supporter" db:"last_supporter"`
	Target        json.RawMessage `json:"target" db:"target"`
	RunningText   json.RawMessage `json:"running_text" db:"running_text"`
	Subathon      json.RawMessage `json:"subathon" db:"subathon"`
	Qrcode        json.RawMessage `json:"qrcode" db:"qrcode"`
	Key           string          `json:"key" db:"key"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	Goals         []Goal
}
