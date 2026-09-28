package models

type EventDetail struct {
	Event      Event
	Organizer  User
	Community  *Community
	Registered int
}
