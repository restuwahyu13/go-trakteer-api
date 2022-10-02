package models

import "time"

type Activities struct {
	Id         int       `json:"id,omitempty" db:"id"`
	UserId     uint      `json:"user_id,omitempty" db:"user_id"`
	UsersInfo  string    `json:"users_info,omitempty" db:"user_info"`
	FirstLogin time.Time `json:"first_login,omitempty" db:"first_login"`
	LastLogin  time.Time `json:"last_login,omitempty" db:"last_login"`
	Users      *Users    `json:"users,omitempty"`
}
