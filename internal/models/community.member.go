package models

import "time"

type CommunityMember struct {
	MemberID    uint      `json:"member_id"`
	UserID      uint      `json:"user_id"`
	CommunityID uint      `json:"community_id"`
	JoinedAt    time.Time `json:"joined_at"`
}
