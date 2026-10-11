package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

var ErrInvalidEventData = errors.New("invvalid event data")

func (s *EventService) CreateEvent(userID uint, req dto.CreateEventRequest) (uint, error) {
	req.Title = strings.TrimSpace(req.Title)
	req.Category = strings.TrimSpace(req.Category)

	if req.Title == "" || req.Category == "" {
		return 0, ErrInvalidEventData
	}

	if _, err := time.Parse("2006-01-02", req.EventDate); err != nil {
		return 0, fmt.Errorf("%w: invalid event date", ErrInvalidEventData)
	}

	start, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return 0, fmt.Errorf("%w: invalid start time", ErrInvalidEventData)
	}

	end, err := time.Parse("15:04", req.EndTime)
	if err != nil || !end.After(start) {
		return 0, fmt.Errorf("%w: end time must be after start time", ErrInvalidEventData)
	}

	for _, speaker := range req.Speakers {
		if strings.TrimSpace(speaker) == "" || len([]rune(speaker)) > 100 {
			return 0, fmt.Errorf("%w: invalid speaker name", ErrInvalidEventData)
		}
	}
	return s.EventRepo.CreateEvent(userID, req)
}
