package models

import (
	"time"
)

type UsersSocialLink struct {
	FacebookLink  string `db:"facebook_link"`
	YoutubeLink   string `db:"youtube_link"`
	InstagramLink string `db:"instagram_link"`
	LinkendinLink string `db:"linkedin_link"`
	GithubLink    string `db:"github_link"`
	DribbleLink   string `db:"dribble_link"`
}

type Users struct {
	Id          uint        `json:"id" db:"id"`
	Username    string      `json:"username" db:"username"`
	Name        string      `json:"name" db:"name"`
	Email       string      `json:"email" db:"email"`
	Password    string      `json:"password,omitempty" db:"password"`
	Active      bool        `json:"active" db:"active"`
	Verified    bool        `json:"verified" db:"verified"`
	SocialLink  *string     `json:"social_link,omitempty" db:"social_link"`
	VideoLink   *string     `json:"video_link,omitempty" db:"video_link"`
	Banner      *string     `json:"banner,omitempty" db:"banner"`
	Photo       *string     `json:"photo,omitempty" db:"photo"`
	RoleId      uint        `json:"role_id,omitempty" db:"role_id"`
	CategorieId uint        `json:"categorie_id,omitempty" db:"categorie_id"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at" db:"deleted_at"`
	Role        *Roles      `json:"role,omitempty"`
	Categorie   *Categories `json:"categorie,omitempty"`
}
