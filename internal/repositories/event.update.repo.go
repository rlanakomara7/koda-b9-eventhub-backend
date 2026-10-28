package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

var (
	ErrUpdateEventNotFound  = errors.New("event not found")
	ErrUpdateEventForbidden = errors.New("you cannot edit this event")
)

func (r *EventRepository) UpdateEvent(
	userID uint,
	eventID uint,
	req dto.UpdateEventRequest,
) error {
	ctx := context.Background()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var role string

	err = tx.QueryRow(ctx, `
		SELECT r.role_name
		FROM users u
		JOIN roles r ON u.role_id = r.role_id
		WHERE u.user_id = $1
		  AND u.status = 'active'
	`, userID).Scan(&role)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotOrganizer
	}
	if err != nil {
		return err
	}
	if role != "organizer" {
		return ErrNotOrganizer
	}

	var ownerID uint
	var oldBannerURL string

	err = tx.QueryRow(ctx, `
		SELECT user_id, COALESCE(banner_url, '')
		FROM events
		WHERE event_id = $1
		FOR UPDATE
	`, eventID).Scan(&ownerID, &oldBannerURL)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrUpdateEventNotFound
	}
	if err != nil {
		return err
	}

	if ownerID != userID {
		return ErrUpdateEventForbidden
	}

	var categoryID uint

	err = tx.QueryRow(ctx, `
		SELECT category_id
		FROM event_categories
		WHERE category_name = $1
	`, req.Category).Scan(&categoryID)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrCategoryNotFound
	}
	if err != nil {
		return err
	}

	var communityID any

	if req.CommunityID != nil {
		if *req.CommunityID == 0 {
			return ErrCommunityNotFound
		}

		var id uint

		err = tx.QueryRow(ctx, `
			SELECT community_id
			FROM communities
			WHERE community_id = $1
			  AND status = 'active'
		`, *req.CommunityID).Scan(&id)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrCommunityNotFound
		}
		if err != nil {
			return err
		}

		communityID = id
	}

	bannerURL := oldBannerURL

	if req.BannerURL != "" {
		bannerURL = req.BannerURL
	}

	_, err = tx.Exec(ctx, `
		UPDATE events
		SET
			community_id = $2,
			title = $3,
			description = $4,
			banner_url = $5,
			event_date = $6,
			start_time = $7,
			end_time = $8,
			format = $9,
			location = $10,
			capacity = $11
		WHERE event_id = $1
	`,
		eventID,
		communityID,
		req.Title,
		req.Description,
		bannerURL,
		req.EventDate,
		req.StartTime,
		req.EndTime,
		req.Format,
		req.Location,
		req.Capacity,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		DELETE FROM event_category_mappings
		WHERE event_id = $1
	`, eventID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO event_category_mappings
			(event_id, category_id)
		VALUES ($1, $2)
	`, eventID, categoryID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		DELETE FROM event_speakers
		WHERE event_id = $1
	`, eventID)
	if err != nil {
		return err
	}

	for _, speaker := range req.Speakers {
		_, err = tx.Exec(ctx, `
			INSERT INTO event_speakers (event_id, name)
			VALUES ($1, $2)
		`, eventID, speaker)

		if err != nil {
			return err
		}
	}

	//notifications
	_, err = tx.Exec(ctx, `
    INSERT INTO notifications (
        user_id, type, title, message, path
    )
    SELECT DISTINCT
        em.user_id,
        'update',
        'Event updated',
        'An event you joined has been updated: ' || e.title,
        '/event/' || e.event_id::text
    FROM event_members em
    JOIN events e ON e.event_id = em.event_id
    WHERE em.event_id = $1
      AND em.user_id <> $2
`, eventID, userID)

	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	if r.Redis != nil {
		keys := []string{
			fmt.Sprintf("eventhub:chace:event:detail:%d", eventID),
			fmt.Sprintf("event:%d", eventID),
		}

		if err := r.Redis.Del(ctx, keys...).Err(); err != nil {
			log.Printf(
				"failed to invalidate event cache %d: %v",
				eventID,
				err,
			)
		}
	}

	return nil
}
