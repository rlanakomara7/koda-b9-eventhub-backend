package dto

type UpdateEventRequest = CreateEventRequest

type UpdateEventResponse struct {
	Message string `json:"message"`
	EventID uint   `json:"event_id"`
}
