package repositories

import (
	"context"
	"time"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *EventRepository) CountEventsPaginated(
	ctx context.Context,
	search string,
	format string,
	category string,
) (int, error) {
	const query = `
		SELECT COUNT(*)
		FROM events e
		WHERE e.status = 'active'
		  AND ($1 = '' OR e.title ILIKE '%' || $1 || '%')
		  AND ($2 = '' OR e.format::text = $2)
		  AND (
		    $3 = '' OR EXISTS (
		      SELECT 1
		      FROM event_category_mappings ecm
		      JOIN event_categories ec
		        ON ec.category_id = ecm.category_id
		      WHERE ecm.event_id = e.event_id
		        AND LOWER(ec.category_name) = LOWER($3)
		    )
		  )
	`

	var total int

	err := r.DB.QueryRow(
		ctx,
		query,
		search,
		format,
		category,
	).Scan(&total)

	return total, err
}

func (r *EventRepository) GetEventsPage(
	ctx context.Context,
	search string,
	format string,
	category string,
	page int,
	limit int,
) ([]dto.EventCardResponse, error) {
	const query = `
		SELECT
			e.event_id,
			e.title,
			COALESCE(e.banner_url, ''),
			e.event_date,
			e.start_time,
			e.format::text,
			COALESCE(e.location, ''),
			e.capacity,
			COALESCE(
				(
					SELECT ARRAY_AGG(
						ec.category_name
						ORDER BY ec.category_name
					)
					FROM event_category_mappings ecm
					JOIN event_categories ec
						ON ec.category_id = ecm.category_id
					WHERE ecm.event_id = e.event_id
				),
				ARRAY[]::text[]
			),
			(
				SELECT COUNT(*)::int
				FROM event_members em
				WHERE em.event_id = e.event_id
			)
		FROM events e
		WHERE e.status = 'active'
		  AND ($1 = '' OR e.title ILIKE '%' || $1 || '%')
		  AND ($2 = '' OR e.format::text = $2)
		  AND (
		    $3 = '' OR EXISTS (
		      SELECT 1
		      FROM event_category_mappings ecm
		      JOIN event_categories ec
		        ON ec.category_id = ecm.category_id
		      WHERE ecm.event_id = e.event_id
		        AND LOWER(ec.category_name) = LOWER($3)
		    )
		  )
		ORDER BY
			e.event_date ASC,
			e.start_time ASC,
			e.event_id ASC
		LIMIT $4 OFFSET $5
	`

	offset := (page - 1) * limit

	rows, err := r.DB.Query(
		ctx,
		query,
		search,
		format,
		category,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := make([]dto.EventCardResponse, 0)

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

		if len(startTime) >= 5 {
			event.StartTime = startTime[:5]
		} else {
			event.StartTime = startTime
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
