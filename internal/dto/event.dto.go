package dto

type OrganizerResponse struct {
	UserID uint `json:"user_id"`

	Name string `json:"name"`

	Email string `json:"email"`
}

type CommunityResponse struct {
	CommunityID uint `json:"community_id"`

	Name string `json:"name"`
}

type EventDetailResponse struct {
	EventID         uint               `json:"event_id"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	BannerURL       string             `json:"banner_url"`
	EventDate       string             `json:"event_date"`
	StartTime       string             `json:"start_time"`
	EndTime         string             `json:"end_time"`
	Format          string             `json:"format"`
	Location        string             `json:"location"`
	Capacity        int                `json:"capacity"`
	Status          string             `json:"status"`
	Organizer       OrganizerResponse  `json:"organizer"`
	Community       *CommunityResponse `json:"community"`
	TotalRegistered int                `json:"total_registered"`
}

type EventInfoResponse struct {
	Date string `json:"date"`

	Time string `json:"time"`

	Location string `json:"location"`

	Capacity int `json:"capacity"`

	Registered int `json:"registered"`

	SpotsLeft int `json:"spots_left"`
}
