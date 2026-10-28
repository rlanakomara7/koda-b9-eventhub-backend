package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

type NotificationRepository struct {
	DB *pgxpool.Pool
}

func NewNotificationRepository(db *pgxpool.Pool) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) GetNotifications(
	ctx context.Context,
	userID uint,
) ([]dto.NotificationResponse, error) {

	rows, err := r.DB.Query(ctx, `
        SELECT
            notification_id,
            type,
            title,
            message,
            path,
            is_read,
            created_at
        FROM notifications
        WHERE user_id = $1
        ORDER BY created_at DESC, notification_id DESC
        LIMIT 100
    `, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := []dto.NotificationResponse{}

	for rows.Next() {
		var item dto.NotificationResponse

		err := rows.Scan(
			&item.ID,
			&item.Type,
			&item.Title,
			&item.Message,
			&item.Path,
			&item.Read,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		notifications = append(notifications, item)
	}

	return notifications, rows.Err()
}

func (r *NotificationRepository) MarkRead(
	ctx context.Context,
	userID uint,
	notificationID uint,
) error {

	var id uint

	return r.DB.QueryRow(ctx, `
        UPDATE notifications
        SET is_read = TRUE
        WHERE notification_id = $1
          AND user_id = $2
        RETURNING notification_id
    `, notificationID, userID).Scan(&id)
}

func (r *NotificationRepository) MarkAllRead(
	ctx context.Context,
	userID uint,
) error {

	_, err := r.DB.Exec(ctx, `
        UPDATE notifications
        SET is_read = TRUE
        WHERE user_id = $1
          AND is_read = FALSE
    `, userID)

	return err
}
