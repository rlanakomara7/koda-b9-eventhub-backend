package dto

type DeleteEventResponse struct {
	Message string `json:"message"`
	EventID uint   `json:"event_id"`
}
