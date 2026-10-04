package services

import (
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type EventService struct {
	EventRepo *repositories.EventRepository
}

func NewEventService(
	eventRepo *repositories.EventRepository,
) *EventService {

	return &EventService{
		EventRepo: eventRepo,
	}
}

// get events
func (s *EventService) GetEvents(
	search string,
	format string,
) ([]models.Event, error) {

	events, err := s.EventRepo.GetEvents(
		search,
		format,
	)

	if err != nil {
		return nil, err
	}
	return events, nil
}

// event detail
func (s *EventService) GetEventDetail(
	id uint,
) (*models.Event, error) {

	event, err := s.EventRepo.GetEventByID(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

// upcoming events
func (s *EventService) GetUpcomingEvents() ([]models.Event, error) {

	return s.EventRepo.GetUpcomingEvents()
}

// get myevents
func (s *EventService) GetMyEvents(UserID uint) ([]models.Event, error) {

	return s.EventRepo.GetMyEvents(UserID)
}
