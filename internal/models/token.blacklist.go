package models

import "time"

type TokenBlackList struct {
	ID        uint      `json:"id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
