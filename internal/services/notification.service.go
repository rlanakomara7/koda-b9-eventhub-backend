package services

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type NotificationService struct {
	Repo *repositories.NotificationRepository
}

func NewNotificationService(
	repo *repositories.NotificationRepository,
) *NotificationService {
	return &NotificationService{Repo: repo}
}

func (s *NotificationService) GetNotifications(
	ctx context.Context,
	userID uint,
) ([]dto.NotificationResponse, error) {
	return s.Repo.GetNotifications(ctx, userID)
}

func (s *NotificationService) MarkRead(
	ctx context.Context,
	userID uint,
	notificationID uint,
) error {
	return s.Repo.MarkRead(ctx, userID, notificationID)
}

func (s *NotificationService) MarkAllRead(
	ctx context.Context,
	userID uint,
) error {
	return s.Repo.MarkAllRead(ctx, userID)
}
