package dto

type CommunityResponse struct {
	CommunityID   uint     `json:"community_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	BannerURL     string   `json:"banner_url"`
	Categories    []string `json:"categories"`
	MemberCount   int      `json:"member_count"`
	UpcomingCount int      `json:"upcoming_count"`
	IsJoined      bool     `json:"is_joined"`
}

type CommunityListResponse struct {
	Data []CommunityResponse `json:"data"`
}
