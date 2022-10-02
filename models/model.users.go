package models

import (
	"time"
)

type Users struct {
	Id        int        `json:"id,omitempty" db:"id"`
	Name      string     `json:"name,omitempty" db:"name"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Active    bool       `json:"active,omitempty" db:"active"`
	Verified  bool       `json:"verified,omitempty" db:"verified"`
	Photo     *string    `json:"photo,omitempty" db:"photo"`
	RoleId    uint       `json:"role_id,omitempty" db:"role_id"`
	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Role      *Roles     `json:"role,omitempty"`
}
