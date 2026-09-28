package models

import "time"

type Event struct {
	EventID     uint      `json:"event_id"`
	UserID      uint      `json:"user_id"`
	CommunityID *uint     `json:"community_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BannerURL   string    `json:"banner_url"`
	EventDate   time.Time `json:"event_date"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	Format      string    `json:"format"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
