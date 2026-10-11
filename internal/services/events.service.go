package services

import (
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
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
) ([]dto.EventCardResponse, error) {

	return s.EventRepo.GetEvents(
		search,
		format,
	)
}

// event detail
func (s *EventService) GetEventDetail(id uint) (*dto.EventDetailResponse, error) {

	return s.EventRepo.GetEventDetail(id)
}

// upcoming events
func (s *EventService) GetUpcomingEvents() ([]models.Event, error) {

	return s.EventRepo.GetUpcomingEvents()
}

// get myevents
func (s *EventService) GetMyEvents(UserID uint) ([]dto.MyEventResponse, error) {

	return s.EventRepo.GetMyEvents(UserID)
}
