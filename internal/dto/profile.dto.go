package dto

type ProfileResponse struct {
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	AvatarURL string       `json:"avatar_url"`
	Location  string       `json:"location"`
	JoinedAt  string       `json:"joined_at"`
	Role      string       `json:"role"`
	Bio       string       `json:"bio"`
	Stats     ProfileStats `json:"stats"`
}
type ProfileStats struct {
	Events      int `json:"events"`
	Communities int `json:"communities"`
	Saved       int `json:"saved"`
}

type ProfileSuccessResponse struct {
	Message string          `json:"message"`
	Data    ProfileResponse `json:"data"`
}

type UpdateProfileRequest struct {
	Name      string `json:"name" form:"name" binding:"required,min=2,max=100"`
	AvatarURL string `json:"avatar_url" form:"avatar_url"`
	Location  string `json:"location" form:"location" binding:"omitempty,max=150"`
	Bio       string `json:"bio" form:"bio" binding:"omitempty,max=500"`
}
