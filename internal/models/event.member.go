package models

import "time"

type EventMember struct {
	EventMemberID uint      `json:"event_member_id"`
	EventID       uint      `json:"event_id"`
	UserID        uint      `json:"user_id"`
	JoinedAt      time.Time `json:"joined_at"`
}
