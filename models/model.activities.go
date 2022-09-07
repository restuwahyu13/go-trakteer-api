package models

import "time"

type Activities struct {
	Id         int       `db:"id"`
	UserId     int       `db:"user_id"`
	UsersInfo  string    `db:"user_info"`
	FirstLogin time.Time `db:"first_login"`
	LastLogin  time.Time `db:"last_login"`
	Users      Users
}
