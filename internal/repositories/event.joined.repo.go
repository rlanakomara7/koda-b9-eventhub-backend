package repositories

import (
	"context"
	"time"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *EventMemberRepository) GetJoinedEvents(
	userID uint,
) ([]dto.EventCardResponse, error) {

	query := `
		SELECT
			e.event_id,
			e.title,
			COALESCE(e.banner_url, ''),
			e.event_date,
			e.start_time,
			e.format::text,
			COALESCE(e.location, ''),
			e.capacity,

			ARRAY(
				SELECT ec.category_name
				FROM event_category_mappings ecm
				JOIN event_categories ec
					ON ec.category_id = ecm.category_id
				WHERE ecm.event_id = e.event_id
			),

			(
				SELECT COUNT(*)::int
				FROM event_members em
				WHERE em.event_id = e.event_id
			)

		FROM events e
		WHERE EXISTS (
			SELECT 1
			FROM event_members em
			WHERE em.event_id = e.event_id
			AND em.user_id = $1
		)

		ORDER BY e.event_date ASC
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []dto.EventCardResponse{}

	for rows.Next() {
		var event dto.EventCardResponse
		var eventDate time.Time
		var startTime string

		err := rows.Scan(
			&event.EventID,
			&event.Title,
			&event.BannerURL,
			&eventDate,
			&startTime,
			&event.Format,
			&event.Location,
			&event.Capacity,
			&event.Categories,
			&event.AttendeeCount,
		)

		if err != nil {
			return nil, err
		}

		event.EventDate = eventDate.Format("2006-01-02")
		event.StartTime = startTime

		if len(startTime) >= 5 {
			event.StartTime = startTime[:5]
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
