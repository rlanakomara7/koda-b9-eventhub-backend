package dto

import "github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"

type OrganizerResponse struct {
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type EventCommunityResponse struct {
	CommunityID uint   `json:"community_id"`
	Name        string `json:"name"`
}

type EventDetailResponse struct {
	EventID     uint   `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	BannerURL   string `json:"banner_url"`

	Categories   []string `json:"categories"`
	Availability string   `json:"availability"`

	EventDate string `json:"event_date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Format    string `json:"format"`
	Location  string `json:"location"`

	Registered int `json:"registered"`
	Capacity   int `json:"capacity"`
	SpotsLeft  int `json:"spots_left"`

	IsJoined bool `json:"is_joined"`
	IsSaved  bool `json:"is_saved"`

	Organizer OrganizerResponse       `json:"organizer"`
	Community *EventCommunityResponse `json:"community"`

	Speakers []EventSpeakerResponse `json:"speakers"`

	DiscussionCount int                       `json:"discussion_count"`
	Discussions     []EventDiscussionResponse `json:"discussions"`
}

type EventInfoResponse struct {
	Date       string `json:"date"`
	Time       string `json:"time"`
	Location   string `json:"location"`
	Capacity   int    `json:"capacity"`
	Registered int    `json:"registered"`
	SpotsLeft  int    `json:"spots_left"`
}

type EventListSuccessResponse struct {
	Data []models.Event `json:"data"`
}

type EventListNotFoundResponse struct {
	Data    []models.Event `json:"data"`
	Message string         `json:"message"`
}

type EventDetailSuccessResponse struct {
	Data *EventDetailResponse `json:"data"`
}

// card event
type EventCardResponse struct {
	EventID       uint     `json:"event_id"`
	Title         string   `json:"title"`
	BannerURL     string   `json:"banner_url"`
	Categories    []string `json:"categories"`
	EventDate     string   `json:"event_date"`
	StartTime     string   `json:"start_time"`
	Format        string   `json:"format"`
	Location      string   `json:"location"`
	AttendeeCount int      `json:"attendee_count"`
	Capacity      int      `json:"capacity"`
}

type EventCardListResponse struct {
	Data []EventCardResponse `json:"data"`
}

type EventCardNotFoundResponse struct {
	Data    []EventCardResponse `json:"data"`
	Message string              `json:"message"`
}

type EventSpeakerResponse struct {
	SpeakerID uint   `json:"speaker_id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	Company   string `json:"company"`
	AvatarURL string `json:"avatar_url"`
}

type EventDiscussionResponse struct {
	DiscussionID uint   `json:"discussion_id"`
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	AvatarURL    string `json:"avatar_url"`
	Message      string `json:"message"`
	CreatedAt    string `json:"created_at"`
}

type MyEventResponse struct {
	EventID     uint   `json:"event_id"`
	CommunityID uint   `json:"community_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	BannerURL   string `json:"banner_url"`
	EventDate   string `json:"event_date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Format      string `json:"format"`
	Location    string `json:"location"`
	Capacity    int    `json:"capacity"`
	Status      string `json:"status"`
}

type MyEventListResponse struct {
	Data []MyEventResponse `json:"data"`
}
