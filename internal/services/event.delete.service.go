package services

import (
	"context"
	"fmt"
)

func (s *EventService) DeleteEvent(
	ctx context.Context,
	userID uint,
	eventID uint,
) error {
	if userID == 0 || eventID == 0 {
		return fmt.Errorf("%w: invalid user or event ID", ErrInvalidEventData)
	}

	return s.EventRepo.DeleteEvent(ctx, userID, eventID)
}
