package models

import "time"

type Activities struct {
	Id         uint      `db:"id"`
	UserId     uint      `db:"user_id"`
	UsersInfo  string    `db:"user_info"`
	FirstLogin time.Time `db:"first_login"`
	LastLogin  time.Time `db:"last_login"`
	Users      Users
}
