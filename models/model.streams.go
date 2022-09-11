package models

import (
	"encoding/json"
	"time"
)

type Streams struct {
	Id            uint            `db:"id"`
	Notification  json.RawMessage `db:"notification"`
	Leaderboard   json.RawMessage `db:"leaderboard"`
	LastSupporter json.RawMessage `db:"last_supporter"`
	Target        json.RawMessage `db:"target"`
	RunningText   json.RawMessage `db:"running_text"`
	Subathon      json.RawMessage `db:"subathon"`
	Qrcode        json.RawMessage `db:"qrcode"`
	Key           string          `db:"key"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
	Goals         []Goal
}
