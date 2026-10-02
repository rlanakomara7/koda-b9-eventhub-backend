package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
)

type EventRepository struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func NewEventRepository(db *pgxpool.Pool, redisClient *redis.Client) *EventRepository {
	return &EventRepository{
		DB:    db,
		Redis: redisClient,
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

	ctx := context.Background()
	key := fmt.Sprintf("event:%d", id)

	//cache redis
	cache, err := r.Redis.Get(ctx, key).Result()

	if err == nil {
		fmt.Println("REDIS DATA FOUND:", key)

		var event models.Event

		err = json.Unmarshal([]byte(cache), &event)
		if err != nil {
			fmt.Println("JSON UNMARSHAL ERROR:", err)
		} else {
			fmt.Println("CACHE HIT:", key)
			return &event, nil
		}
	}

	if err != nil && err != redis.Nil {
		fmt.Println("REDIS GET ERROR:", err)
	}

	fmt.Println("CACHE MISS", key)
	// QUERY DATABASE
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

	err = r.DB.QueryRow(
		ctx,
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

	//struct to json
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = r.Redis.Set(
		ctx,
		key,
		data,
		10*time.Minute,
	).Err()

	if err != nil {
		fmt.Println("REDIS SET ERROR:", err)
	} else {
		fmt.Println("CACHE SAVED:", key)
	}
	return event, nil
}

func (r *EventRepository) GetEventDetail(
	id uint,
) (*models.EventDetail, error) {

	query := `
	SELECT

	e.event_id,
	e.title,
	e.description,
	e.banner_url,

	e.event_date,
	e.start_time,
	e.end_time,

	e.format,
	e.location,
	e.capacity,
	e.status,


	u.user_id,
	u.name,
	u.email,


	c.community_id,
	c.name,


	(
		SELECT COUNT(*)
		FROM event_members em
		WHERE em.event_id=e.event_id
	)


	FROM events e


	JOIN users u
	ON e.user_id=u.user_id


	LEFT JOIN communities c
	ON e.community_id=c.community_id


	WHERE e.event_id=$1
	`

	detail := &models.EventDetail{}

	var communityID sql.NullInt64
	var communityName sql.NullString

	err := r.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(

		&detail.Event.EventID,
		&detail.Event.Title,
		&detail.Event.Description,
		&detail.Event.BannerURL,

		&detail.Event.EventDate,
		&detail.Event.StartTime,
		&detail.Event.EndTime,

		&detail.Event.Format,
		&detail.Event.Location,
		&detail.Event.Capacity,
		&detail.Event.Status,

		&detail.Organizer.UserID,
		&detail.Organizer.Name,
		&detail.Organizer.Email,

		&communityID,
		&communityName,

		&detail.Registered,
	)

	if err != nil {
		return nil, err
	}

	if communityID.Valid {

		detail.Community = &models.Community{
			CommunityID: uint(communityID.Int64),
			Name:        communityName.String,
		}

	}

	return detail, nil
}
