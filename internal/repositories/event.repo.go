package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
)

type EventRepository struct {
	DB *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		DB: db,
	}
}

// get eventlist
func (r *EventRepository) GetEvents(search string, format string) ([]models.Event, error) {
	query := `SELECT 
				event_id,
				user_id,
		community_id,
		title,
		description,
		banner_url,
		event_date,
		start_time,
		end_time,
		format,
		location,
		capacity,
		status,
		created_at,
		updated_at
		
		FROM events
		WHERE
		($1 = '' OR title ILIKE '%'||$1||'%')
		AND
		($2 = '' OR format::text = $2)
		ORDER BY event_date ASC`

	rows, err := r.DB.Query(
		context.Background(),
		query,
		search,
		format,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.EventID,
			&event.UserID,
			&event.CommunityID,
			&event.Title,
			&event.Description,
			&event.BannerURL,
			&event.EventDate,
			&event.StartTime,
			&event.EndTime,
			&event.Format,
			&event.Location,
			&event.Capacity,
			&event.Status,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

// get event by id
func (r *EventRepository) GetEventByID(id uint) (*models.Event, error) {

	query := `
	SELECT
		event_id,
		user_id,
		community_id,
		title,
		description,
		banner_url,
		event_date,
		start_time,
		end_time,
		format,
		location,
		capacity,
		status,
		created_at,
		updated_at

	FROM events
	WHERE event_id=$1
	`

	event := &models.Event{}

	err := r.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&event.EventID,
		&event.UserID,
		&event.CommunityID,
		&event.Title,
		&event.Description,
		&event.BannerURL,
		&event.EventDate,
		&event.StartTime,
		&event.EndTime,
		&event.Format,
		&event.Location,
		&event.Capacity,
		&event.Status,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return event, nil
}
