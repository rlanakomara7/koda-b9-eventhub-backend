package repositories

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *EventMemberRepository) SaveEvent(eventID, userID uint) error {
	_, err := r.DB.Exec(
		context.Background(),
		`INSERT INTO saved_events (event_id, user_id)
		 VALUES ($1, $2)
		 ON CONFLICT (user_id, event_id) DO NOTHING`,
		eventID, userID,
	)
	return err
}

func (r *EventMemberRepository) UnsaveEvent(eventID, userID uint) error {
	_, err := r.DB.Exec(
		context.Background(),
		`DELETE FROM saved_events
		 WHERE event_id = $1 AND user_id = $2`,
		eventID, userID,
	)
	return err
}

func (r *EventMemberRepository) GetSavedEvents(userID uint) ([]dto.EventCardResponse, error) {
	rows, err := r.DB.Query(
		context.Background(),
		`
		SELECT
			e.event_id,
			e.title,
			COALESCE(e.banner_url, ''),
			e.event_date::text,
			LEFT(e.start_time::text, 5),
			e.format::text,
			COALESCE(e.location, ''),
			e.capacity,
			(SELECT COUNT(*) FROM event_members m
			 WHERE m.event_id = e.event_id),
			ARRAY(
				SELECT c.category_name
				FROM event_category_mappings cm
				JOIN event_categories c ON c.category_id = cm.category_id
				WHERE cm.event_id = e.event_id
			)
		FROM saved_events s
		JOIN events e ON e.event_id = s.event_id
		WHERE s.user_id = $1
		ORDER BY s.saved_event_id DESC
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []dto.EventCardResponse{}

	for rows.Next() {
		var event dto.EventCardResponse

		err := rows.Scan(
			&event.EventID,
			&event.Title,
			&event.BannerURL,
			&event.EventDate,
			&event.StartTime,
			&event.Format,
			&event.Location,
			&event.Capacity,
			&event.AttendeeCount,
			&event.Categories,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, rows.Err()
}
