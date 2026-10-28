package dto

import "time"

type CommunityDiscussionResponse struct {
	DiscussionID uint      `json:"discussion_id"`
	CommunityID  uint      `json:"community_id"`
	UserID       uint      `json:"user_id"`
	Name         string    `json:"name"`
	AvatarURL    string    `json:"avatar_url"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
}
