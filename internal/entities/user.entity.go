package entities

import "time"

type User struct {
	ID              uint       `json:"id"`
	Firstname       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Email           string     `json:"email"`
	Password        string     `json:"-"`
	Username        string     `json:"username"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	AvatarKey       string     `json:"avatar"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
