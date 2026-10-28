package services

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (s *EventMemberService) GetEventDiscussions(
	ctx context.Context,
	eventID uint,
) ([]dto.EventDiscussionResponse, error) {

	return s.EventMemberRepo.GetEventDiscussions(ctx, eventID)
}

func (s *EventMemberService) CreateEventDiscussion(
	ctx context.Context,
	eventID uint,
	userID uint,
	message string,
) (dto.EventDiscussionResponse, error) {

	return s.EventMemberRepo.CreateEventDiscussion(
		ctx,
		eventID,
		userID,
		message,
	)
}
