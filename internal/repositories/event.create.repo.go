package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

var (
	ErrNotOrganizer      = errors.New("only organizer can create events")
	ErrCategoryNotFound  = errors.New("event category not found")
	ErrCommunityNotFound = errors.New("community not found")
)

func (r *EventRepository) CreateEvent(userID uint, req dto.CreateEventRequest) (uint, error) {
	ctx := context.Background()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback(ctx)

	var role string

	err = tx.QueryRow(ctx, `
		SELECT r.role_name
		FROM users u
		JOIN roles r ON u.role_id = r.role_id
		WHERE u.user_id = $1 AND u.status = 'active'
	`, userID).Scan(&role)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, ErrNotOrganizer
	}

	if err != nil {
		return 0, err
	}

	if role != "organizer" {
		return 0, ErrNotOrganizer
	}

	var categoryID uint
	err = tx.QueryRow(ctx, `
	SELECT category_id
	FROM event_categories
	WHERE category_name = $1
	`, req.Category).Scan(&categoryID)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, ErrCategoryNotFound
	}
	if err != nil {
		return 0, err
	}

	var communityID any

	if req.CommunityID != nil {
		if *req.CommunityID == 0 {
			return 0, ErrCommunityNotFound
		}

		var id uint

		err = tx.QueryRow(ctx, `
		SELECT community_id
			FROM communities
			WHERE community_id = $1 AND status = 'active'
		`, *req.CommunityID).Scan(&id)

		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrCommunityNotFound
		}

		if err != nil {
			return 0, err
		}

		communityID = id
	}

	var eventID uint

	err = tx.QueryRow(ctx, `
		INSERT INTO events (
			user_id, community_id, title, description,
			banner_url, event_date, start_time, end_time,
			format, location, capacity
		)
		VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11
		)
		RETURNING event_id
	`,
		userID,
		communityID,
		req.Title,
		req.Description,
		req.BannerURL,
		req.EventDate,
		req.StartTime,
		req.EndTime,
		req.Format,
		req.Location,
		req.Capacity,
	).Scan(&eventID)

	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(ctx, `
	INSERT INTO event_category_mappings (event_id, category_id)
	VALUES ($1, $2)
	`, eventID, categoryID)

	if err != nil {
		return 0, err
	}

	for _, speaker := range req.Speakers {
		_, err = tx.Exec(ctx, `
		INSERT INTO event_speakers (event_id, name)
		VALUES ($1, $2)
		`, eventID, speaker)

		if err != nil {
			return 0, err
		}

	}
	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return eventID, nil
}
