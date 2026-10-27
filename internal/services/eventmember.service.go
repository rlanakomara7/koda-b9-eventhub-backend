package services

import (
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type EventMemberService struct {
	EventMemberRepo *repositories.EventMemberRepository
}

func NewEventMemberService(
	repo *repositories.EventMemberRepository,
) *EventMemberService {
	return &EventMemberService{
		EventMemberRepo: repo,
	}
}

func (s *EventMemberService) JoinEvent(
	eventID uint,
	userID uint,
) error {
	return s.EventMemberRepo.JoinEvent(eventID, userID)
}

func (s *EventMemberService) LeaveEvent(
	eventID uint,
	userID uint,
) error {
	return s.EventMemberRepo.LeaveEvent(eventID, userID)
}

func (s *EventMemberService) GetJoinedEvents(
	userID uint,
) ([]dto.EventCardResponse, error) {
	return s.EventMemberRepo.GetJoinedEvents(userID)
}
