package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (s *EventService) UpdateEvent(
	userID uint,
	eventID uint,
	req dto.UpdateEventRequest,
) error {
	req.Title = strings.TrimSpace(req.Title)
	req.Category = strings.TrimSpace(req.Category)

	if req.Title == "" || req.Category == "" {
		return fmt.Errorf("%w: title and category are required", ErrInvalidEventData)
	}

	if req.Capacity <= 0 {
		return fmt.Errorf("%w: capacity must be greater than zero", ErrInvalidEventData)
	}

	if req.Format != "online" && req.Format != "in_person" {
		return fmt.Errorf("%w: invalid event format", ErrInvalidEventData)
	}

	if _, err := time.Parse("2006-01-02", req.EventDate); err != nil {
		return fmt.Errorf("%w: invalid event date", ErrInvalidEventData)
	}

	start, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return fmt.Errorf("%w: invalid start time", ErrInvalidEventData)
	}

	end, err := time.Parse("15:04", req.EndTime)
	if err != nil || !end.After(start) {
		return fmt.Errorf("%w: end time must be after start time", ErrInvalidEventData)
	}

	for i, speaker := range req.Speakers {
		speaker = strings.TrimSpace(speaker)

		if speaker == "" || len([]rune(speaker)) > 100 {
			return fmt.Errorf("%w: invalid speaker name", ErrInvalidEventData)
		}

		req.Speakers[i] = speaker
	}

	return s.EventRepo.UpdateEvent(userID, eventID, req)
}
