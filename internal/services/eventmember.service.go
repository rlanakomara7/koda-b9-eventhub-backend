package services

import (
	"errors"

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

	// cek sudah join
	alreadyJoin := s.EventMemberRepo.IsJoined(
		eventID,
		userID,
	)

	if alreadyJoin {
		return errors.New("already joined event")
	}

	err := s.EventMemberRepo.JoinEvent(
		eventID,
		userID,
	)
	return err
}
