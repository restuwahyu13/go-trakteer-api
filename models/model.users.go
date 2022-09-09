package models

import "time"

type rolesRelation struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"created_at"`
}

type categoriesRelation struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

type Users struct {
	ID          uint      `db:"id"`
	Username    string    `db:"username"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	RoleID      int       `db:"role_id"`
	CategorieID int       `db:"categorie_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"created_at"`
	Role        rolesRelation
	Categorie   categoriesRelation
}
