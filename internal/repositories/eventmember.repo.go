package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventMemberRepository struct {
	DB *pgxpool.Pool
}

func NewEventMemberRepository(
	db *pgxpool.Pool,
) *EventMemberRepository {

	return &EventMemberRepository{
		DB: db,
	}
}

// JOIN EVENT

func (r *EventMemberRepository) JoinEvent(
	eventID uint,
	userID uint,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO event_members
		(
			event_id,
			user_id
		)
		VALUES
		($1,$2)
		`,
		eventID,
		userID,
	)
	return err
}

// cek user join
func (r *EventMemberRepository) IsJoined(
	eventID uint,
	userID uint,
) bool {

	var count int

	r.DB.QueryRow(
		context.Background(),
		`
		SELECT COUNT(*)
		FROM event_members
		WHERE event_id=$1
		AND user_id=$2
		`,
		eventID,
		userID,
	).Scan(&count)

	return count > 0
}
