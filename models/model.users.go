package models

import (
	"time"
)

type Users struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Active    bool       `json:"active" db:"active"`
	Verified  bool       `json:"verified" db:"verified"`
	Photo     *string    `json:"photo,omitempty" db:"photo"`
	RoleId    uint       `json:"role_id,omitempty" db:"role_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	Role      *Roles     `json:"role,omitempty"`
}
