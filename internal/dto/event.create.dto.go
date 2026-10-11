package dto

type CreateEventRequest struct {
	Title       string   `form:"title" binding:"required,max=150"`
	Description string   `form:"description"`
	Category    string   `form:"category" binding:"required"`
	CommunityID *uint    `form:"community_id"`
	BannerURL   string   `json:"banner_url"`
	EventDate   string   `form:"event_date" binding:"required"`
	StartTime   string   `form:"start_time" binding:"required"`
	EndTime     string   `form:"end_time" binding:"required"`
	Format      string   `form:"format" binding:"required,oneof=online in_person"`
	Location    string   `form:"location"`
	Capacity    int      `form:"capacity" binding:"required,gt=0"`
	Speakers    []string `json:"speakers" form:"-"`
}

type CreateEventResponse struct {
	Message string `json:"message"`
	EventID uint   `json:"event_id"`
}
