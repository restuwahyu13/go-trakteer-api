package models

import "time"

type Activities struct {
	Id         int       `json:"id" db:"id"`
	UserId     uint      `json:"user_id" db:"user_id"`
	UsersInfo  string    `json:"users_info" db:"user_info"`
	FirstLogin time.Time `json:"first_login" db:"first_login"`
	LastLogin  time.Time `json:"last_login" db:"last_login"`
	Users      *Users
}
