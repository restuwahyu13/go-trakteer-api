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
	Id          int         `json:"id" db:"id"`
	Username    string      `json:"username" db:"username"`
	Name        string      `json:"name" db:"name"`
	Email       string      `json:"email,omitempty" db:"email"`
	Password    string      `json:"password,omitempty" db:"password"`
	Active      bool        `json:"active" db:"active"`
	Verified    bool        `json:"verified" db:"verified"`
	SocialLink  string      `json:"social_link" db:"social_link"`
	VideoLink   *string     `json:"video_link" db:"video_link"`
	Banner      *string     `json:"banner" db:"banner"`
	Photo       *string     `json:"photo" db:"photo"`
	RoleId      uint        `json:"role_id" db:"role_id"`
	CategorieId uint        `json:"categorie_id" db:"categorie_id"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at" db:"deleted_at"`
	Role        *Roles      `json:"role,omitempty"`
	Categorie   *Categories `json:"categorie,omitempty"`
}
