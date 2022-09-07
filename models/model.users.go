package models

import "time"

type Users struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	RoleID      int       `db:"role_id"`
	CategorieID int       `db:"categorie_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"created_at"`
}
