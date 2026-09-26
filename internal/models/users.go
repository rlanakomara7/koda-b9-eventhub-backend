package models

import "time"

type User struct {
	UserID    uint      `json:"user_id"`
	RoleID    uint      `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	AvatarURL string    `json:"avatar_url"`
	Job       string    `json:"job"`
	Location  string    `json:"location"`
	Bio       string    `json:"bio"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
