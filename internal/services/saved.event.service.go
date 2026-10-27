package services

import "github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"

func (s *EventMemberService) SaveEvent(eventID, userID uint) error {
	return s.EventMemberRepo.SaveEvent(eventID, userID)
}

func (s *EventMemberService) UnsaveEvent(eventID, userID uint) error {
	return s.EventMemberRepo.UnsaveEvent(eventID, userID)
}

func (s *EventMemberService) GetSavedEvents(userID uint) ([]dto.EventCardResponse, error) {
	return s.EventMemberRepo.GetSavedEvents(userID)
}
