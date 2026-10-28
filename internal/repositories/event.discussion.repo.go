package repositories

import (
	"context"
	"time"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *EventMemberRepository) GetEventDiscussions(
	ctx context.Context,
	eventID uint,
) ([]dto.EventDiscussionResponse, error) {

	query := `
        SELECT
            d.discussion_id,
            d.event_id,
            d.user_id,
            u.name,
            COALESCE(u.avatar_url, ''),
            d.message,
            d.created_at
        FROM event_discussions d
        JOIN users u ON u.user_id = d.user_id
        WHERE d.event_id = $1
        ORDER BY d.created_at DESC, d.discussion_id DESC
        LIMIT 100
    `

	rows, err := r.DB.Query(ctx, query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discussions := []dto.EventDiscussionResponse{}

	for rows.Next() {
		var discussion dto.EventDiscussionResponse
		var createdAt time.Time

		err := rows.Scan(
			&discussion.DiscussionID,
			&discussion.EventID,
			&discussion.UserID,
			&discussion.Name,
			&discussion.AvatarURL,
			&discussion.Message,
			&createdAt,
		)

		if err != nil {
			return nil, err
		}

		discussion.CreatedAt = createdAt.Format(time.RFC3339)

		discussions = append(discussions, discussion)
	}

	return discussions, rows.Err()
}

func (r *EventMemberRepository) CreateEventDiscussion(
	ctx context.Context,
	eventID uint,
	userID uint,
	message string,
) (dto.EventDiscussionResponse, error) {

	var discussion dto.EventDiscussionResponse
	var createdAt time.Time

	query := `
        WITH inserted AS (
            INSERT INTO event_discussions (
                event_id,
                user_id,
                message
            )
            SELECT e.event_id, $2, $3
            FROM events e
            WHERE e.event_id = $1
              AND e.status IN ('active', 'full')
              AND EXISTS (
                  SELECT 1
                  FROM event_members em
                  WHERE em.event_id = e.event_id
                    AND em.user_id = $2
              )
            RETURNING
                discussion_id,
                event_id,
                user_id,
                message,
                created_at
        )
        SELECT
            d.discussion_id,
            d.event_id,
            d.user_id,
            u.name,
            COALESCE(u.avatar_url, ''),
            d.message,
            d.created_at
        FROM inserted d
        JOIN users u ON u.user_id = d.user_id
    `

	err := r.DB.QueryRow(
		ctx,
		query,
		eventID,
		userID,
		message,
	).Scan(
		&discussion.DiscussionID,
		&discussion.EventID,
		&discussion.UserID,
		&discussion.Name,
		&discussion.AvatarURL,
		&discussion.Message,
		&createdAt,
	)

	if err != nil {
		return discussion, err
	}

	discussion.CreatedAt = createdAt.Format(time.RFC3339)

	return discussion, nil
}
