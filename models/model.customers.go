package models

import (
	"time"
)

type CustomersSocialLink struct {
	FacebookLink  string `json:"facebook_link" db:"facebook_link"`
	YoutubeLink   string `json:"youtube_link" db:"youtube_link"`
	InstagramLink string `json:"instagram_link" db:"instagram_link"`
	LinkendinLink string `json:"linkendin_link" db:"linkedin_link"`
	GithubLink    string `json:"github_link" db:"github_link"`
	DribbleLink   string `json:"dribble_link" db:"dribble_link"`
}

type Customers struct {
	Id          int         `json:"id,omitempty" db:"id"`
	Username    string      `json:"username,omitempty" db:"username"`
	Name        string      `json:"name,omitempty" db:"name"`
	Email       string      `json:"email,omitempty" db:"email"`
	Password    string      `json:"password,omitempty" db:"password"`
	Active      bool        `json:"active,omitempty" db:"active"`
	Verified    bool        `json:"verified,omitempty" db:"verified"`
	SocialLink  string      `json:"social_link,omitempty" db:"social_link"`
	VideoLink   *string     `json:"video_link,omitempty" db:"video_link"`
	Banner      *string     `json:"banner,omitempty" db:"banner"`
	Photo       *string     `json:"photo,omitempty" db:"photo"`
	RoleId      uint        `json:"role_id,omitempty" db:"role_id"`
	CategorieId uint        `json:"categorie_id,omitempty" db:"categorie_id"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
	Role        *Roles      `json:"role,omitempty"`
	Categorie   *Categories `json:"categorie,omitempty"`
}
