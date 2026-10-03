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
