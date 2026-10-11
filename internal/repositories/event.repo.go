package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
)

type EventRepository struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func NewEventRepository(
	db *pgxpool.Pool,
	redisClient *redis.Client,
) *EventRepository {

	return &EventRepository{
		DB:    db,
		Redis: redisClient,
	}
}

// get event list
func (r *EventRepository) GetEvents(
	search string,
	format string,
) ([]dto.EventCardResponse, error) {

	query := `
	SELECT
		event_id,
		title,
		banner_url,
		event_date,
		start_time,
		format,
		location,
		capacity
	FROM events
	WHERE status = 'active'
	AND ($1 = '' OR title ILIKE '%' || $1 || '%')
	AND ($2 = '' OR format::text = $2)
	ORDER BY event_date ASC, start_time ASC
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
		search,
		format,
	)

	if err != nil {
		return nil, err
	}

	events := []dto.EventCardResponse{}

	for rows.Next() {

		var event dto.EventCardResponse

		var bannerURL *string
		var location *string
		var eventDate time.Time
		var startTime string

		err := rows.Scan(
			&event.EventID,
			&event.Title,
			&bannerURL,
			&eventDate,
			&startTime,
			&event.Format,
			&location,
			&event.Capacity,
		)

		if err != nil {
			rows.Close()
			return nil, err
		}

		if bannerURL != nil {
			event.BannerURL = *bannerURL
		}

		if location != nil {
			event.Location = *location
		}

		event.EventDate = eventDate.Format("2006-01-02")

		event.StartTime = startTime

		if len(startTime) >= 5 {
			event.StartTime = startTime[:5]
		}

		event.Categories = []string{}

		events = append(
			events,
			event,
		)
	}

	rows.Close()

	for i := range events {

		categoryRows, err := r.DB.Query(
			context.Background(),
			`
			SELECT ec.category_name
			FROM event_categories ec

			JOIN event_category_mappings ecm
			ON ec.category_id = ecm.category_id

			WHERE ecm.event_id = $1
			`,
			events[i].EventID,
		)

		if err != nil {
			return nil, err
		}

		for categoryRows.Next() {

			var categoryName string

			err := categoryRows.Scan(
				&categoryName,
			)

			if err != nil {
				categoryRows.Close()
				return nil, err
			}

			events[i].Categories = append(
				events[i].Categories,
				categoryName,
			)
		}

		categoryRows.Close()

		err = r.DB.QueryRow(
			context.Background(),
			`
			SELECT COUNT(*)
			FROM event_members
			WHERE event_id = $1
			`,
			events[i].EventID,
		).Scan(
			&events[i].AttendeeCount,
		)

		if err != nil {
			return nil, err
		}
	}

	return events, nil
}

// get event by id
func (r *EventRepository) GetEventByID(
	id uint,
) (*models.Event, error) {

	ctx := context.Background()

	key := fmt.Sprintf(
		"event:%d",
		id,
	)

	cache, err := r.Redis.Get(
		ctx,
		key,
	).Result()

	if err == nil {

		fmt.Println(
			"REDIS DATA FOUND:",
			key,
		)

		var event models.Event

		err = json.Unmarshal(
			[]byte(cache),
			&event,
		)

		if err != nil {

			fmt.Println(
				"JSON UNMARSHAL ERROR:",
				err,
			)

		} else {

			fmt.Println(
				"CACHE HIT:",
				key,
			)

			return &event, nil
		}
	}

	if err != nil && err != redis.Nil {

		fmt.Println(
			"REDIS GET ERROR:",
			err,
		)
	}

	fmt.Println(
		"CACHE MISS:",
		key,
	)

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
	WHERE event_id = $1
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

	data, err := json.Marshal(
		event,
	)

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

		fmt.Println(
			"REDIS SET ERROR:",
			err,
		)

	} else {

		fmt.Println(
			"CACHE SAVED:",
			key,
		)
	}

	return event, nil
}

// get event detail
func (r *EventRepository) GetEventDetail(
	id uint,
) (*dto.EventDetailResponse, error) {

	ctx := context.Background()

	key := fmt.Sprintf(
		"eventhub:chace:event:detail:%d",
		id,
	)

	//redis
	cache, err := r.Redis.Get(
		ctx,
		key,
	).Result()

	if err == nil {

		var detail dto.EventDetailResponse

		err = json.Unmarshal(
			[]byte(cache),
			&detail,
		)

		if err == nil {

			fmt.Println(
				"CACHE HIT:",
				key,
			)

			return &detail, nil
		}
	}

	if err != nil && err != redis.Nil {

		fmt.Println(
			"REDIS GET ERROR:",
			err,
		)
	}

	fmt.Println(
		"CACHE MISS:",
		key,
	)

	// GET MAIN EVENT
	// =========================

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
		u.avatar_url,

		c.community_id,
		c.name

	FROM events e

	JOIN users u
		ON e.user_id = u.user_id

	LEFT JOIN communities c
		ON e.community_id = c.community_id

	WHERE e.event_id = $1
	`

	detail := &dto.EventDetailResponse{}

	var description sql.NullString
	var bannerURL sql.NullString
	var location sql.NullString

	var organizerAvatar sql.NullString

	var communityID sql.NullInt64
	var communityName sql.NullString

	var eventDate time.Time
	var startTime string
	var endTime string

	var eventStatus string

	err = r.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&detail.EventID,
		&detail.Title,
		&description,
		&bannerURL,
		&eventDate,
		&startTime,
		&endTime,
		&detail.Format,
		&location,
		&detail.Capacity,
		&eventStatus,

		&detail.Organizer.UserID,
		&detail.Organizer.Name,
		&organizerAvatar,

		&communityID,
		&communityName,
	)

	if err != nil {
		return nil, err
	}

	if description.Valid {
		detail.Description = description.String
	}

	if bannerURL.Valid {
		detail.BannerURL = bannerURL.String
	}

	if location.Valid {
		detail.Location = location.String
	}

	if organizerAvatar.Valid {
		detail.Organizer.AvatarURL =
			organizerAvatar.String
	}

	detail.EventDate =
		eventDate.Format("2006-01-02")

	detail.StartTime = startTime
	detail.EndTime = endTime

	if len(startTime) >= 5 {
		detail.StartTime =
			startTime[:5]
	}

	if len(endTime) >= 5 {
		detail.EndTime =
			endTime[:5]
	}

	if communityID.Valid {

		detail.Community =
			&dto.EventCommunityResponse{
				CommunityID: uint(
					communityID.Int64,
				),
				Name: communityName.String,
			}
	}

	// REGISTERED / ATTENDEES
	// =========================

	err = r.DB.QueryRow(
		ctx,
		`
		SELECT COUNT(*)
		FROM event_members
		WHERE event_id = $1
		`,
		id,
	).Scan(
		&detail.Registered,
	)

	if err != nil {
		return nil, err
	}

	// SPOTS LEFT
	// =========================

	detail.SpotsLeft =
		detail.Capacity -
			detail.Registered

	if detail.SpotsLeft < 0 {
		detail.SpotsLeft = 0
	}

	// AVAILABILITY

	if eventStatus == "completed" {

		detail.Availability =
			"completed"

	} else if eventStatus == "full" ||
		detail.SpotsLeft == 0 {

		detail.Availability =
			"full"

	} else {

		detail.Availability =
			"available"
	}

	// CATEGORIES
	detail.Categories = []string{}

	categoryRows, err := r.DB.Query(
		ctx,
		`
		SELECT
			ec.category_name
		FROM event_categories ec

		JOIN event_category_mappings ecm
			ON ec.category_id =
			ecm.category_id

		WHERE ecm.event_id = $1

		ORDER BY ec.category_id ASC
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	for categoryRows.Next() {

		var category string

		err = categoryRows.Scan(
			&category,
		)

		if err != nil {
			categoryRows.Close()
			return nil, err
		}

		detail.Categories =
			append(
				detail.Categories,
				category,
			)
	}

	categoryRows.Close()

	// SPEAKERS

	detail.Speakers =
		[]dto.EventSpeakerResponse{}

	speakerRows, err := r.DB.Query(
		ctx,
		`
		SELECT
			speaker_id,
			name,
			job,
			company,
			avatar_url
		FROM event_speakers
		WHERE event_id = $1
		ORDER BY speaker_id ASC
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	for speakerRows.Next() {

		var speaker dto.EventSpeakerResponse

		var job sql.NullString
		var company sql.NullString
		var avatarURL sql.NullString

		err = speakerRows.Scan(
			&speaker.SpeakerID,
			&speaker.Name,
			&job,
			&company,
			&avatarURL,
		)

		if err != nil {
			speakerRows.Close()
			return nil, err
		}

		if job.Valid {
			speaker.Job =
				job.String
		}

		if company.Valid {
			speaker.Company =
				company.String
		}

		if avatarURL.Valid {
			speaker.AvatarURL =
				avatarURL.String
		}

		detail.Speakers =
			append(
				detail.Speakers,
				speaker,
			)
	}

	speakerRows.Close()

	// DISCUSSIONS
	detail.Discussions =
		[]dto.EventDiscussionResponse{}

	discussionRows, err := r.DB.Query(
		ctx,
		`
		SELECT
			ed.discussion_id,
			u.user_id,
			u.name,
			u.avatar_url,
			ed.message,
			ed.created_at

		FROM event_discussions ed

		JOIN users u
			ON ed.user_id =
			u.user_id

		WHERE ed.event_id = $1

		ORDER BY ed.created_at DESC
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	for discussionRows.Next() {

		var discussion dto.EventDiscussionResponse

		var avatarURL sql.NullString
		var createdAt time.Time

		err = discussionRows.Scan(
			&discussion.DiscussionID,
			&discussion.UserID,
			&discussion.Name,
			&avatarURL,
			&discussion.Message,
			&createdAt,
		)

		if err != nil {
			discussionRows.Close()
			return nil, err
		}

		if avatarURL.Valid {
			discussion.AvatarURL =
				avatarURL.String
		}

		discussion.CreatedAt =
			createdAt.Format(
				"2006-01-02 15:04:05",
			)

		detail.Discussions =
			append(
				detail.Discussions,
				discussion,
			)
	}

	discussionRows.Close()

	detail.DiscussionCount =
		len(detail.Discussions)

	// REDIS SAVE
	data, err := json.Marshal(
		detail,
	)

	if err != nil {

		fmt.Println(
			"JSON MARSHAL ERROR:",
			err,
		)

		return detail, nil
	}

	err = r.Redis.Set(
		ctx,
		key,
		data,
		10*time.Minute,
	).Err()

	if err != nil {

		fmt.Println(
			"REDIS SET ERROR:",
			err,
		)

	} else {

		fmt.Println(
			"CACHE SAVED:",
			key,
		)
	}

	return detail, nil
}

// get upcoming events
func (r *EventRepository) GetUpcomingEvents() (
	[]models.Event,
	error,
) {

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
	WHERE event_date >= CURRENT_DATE
	AND status != 'completed'
	ORDER BY event_date ASC, start_time ASC
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []models.Event{}

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

		events = append(
			events,
			event,
		)
	}

	return events, nil
}

// get my event
func (r *EventRepository) GetMyEvents(
	userID uint,
) ([]dto.MyEventResponse, error) {

	query := `
	SELECT
		e.event_id,
		e.community_id,
		e.title,
		e.description,
		e.banner_url,
		e.event_date,
		e.start_time,
		e.end_time,
		e.format,
		e.location,
		e.capacity,
		e.status
	FROM events e

	JOIN event_members em
	ON e.event_id = em.event_id

	WHERE em.user_id = $1

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

	events := []dto.MyEventResponse{}

	for rows.Next() {

		var event dto.MyEventResponse
		var communityID sql.NullInt64
		var description sql.NullString
		var bannerURL sql.NullString
		var location sql.NullString

		var eventDate time.Time
		var startTime string
		var endTime string

		err := rows.Scan(
			&event.EventID,
			&event.CommunityID,
			&event.Title,
			&event.Description,
			&event.BannerURL,
			&eventDate,
			&startTime,
			&endTime,
			&event.Format,
			&event.Location,
			&event.Capacity,
			&event.Status,
		)

		if err != nil {
			return nil, err
		}

		if communityID.Valid {
			event.CommunityID = uint(communityID.Int64)
		}

		if description.Valid {
			event.Description = description.String
		}

		if bannerURL.Valid {
			event.BannerURL = bannerURL.String
		}

		if location.Valid {
			event.Location = location.String
		}

		event.EventDate = eventDate.Format("2006-01-02")

		event.StartTime = startTime
		event.EndTime = endTime

		if len(startTime) >= 5 {
			event.StartTime = startTime[:5]
		}

		if len(endTime) >= 5 {
			event.EndTime = endTime[:5]
		}

		events = append(events, event)
	}

	return events, nil
}
