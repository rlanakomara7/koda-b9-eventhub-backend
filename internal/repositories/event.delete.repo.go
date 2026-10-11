package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	ErrDeleteEventNotFound  = errors.New("event not found")
	ErrDeleteEventForbidden = errors.New(
		"you do not have permission to delete this event",
	)
)

func (r *EventRepository) DeleteEvent(
	ctx context.Context,
	userID uint,
	eventID uint,
) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// ✅ TAMBAH: Hanya organizer aktif
	var role string

	err = tx.QueryRow(ctx, `
		SELECT r.role_name
		FROM users u
		JOIN roles r ON u.role_id = r.role_id
		WHERE u.user_id = $1
		  AND u.status = 'active'
	`, userID).Scan(&role)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrDeleteEventForbidden
	}
	if err != nil {
		return err
	}

	if role != "organizer" {
		return ErrDeleteEventForbidden
	}

	var ownerID uint

	err = tx.QueryRow(ctx, `
		SELECT user_id
		FROM events
		WHERE event_id = $1
		FOR UPDATE
	`, eventID).Scan(&ownerID)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrDeleteEventNotFound
	}
	if err != nil {
		return err
	}

	if ownerID != userID {
		return ErrDeleteEventForbidden
	}

	queries := []string{
		`DELETE FROM event_members WHERE event_id = $1`,
		`DELETE FROM saved_events WHERE event_id = $1`,
		`DELETE FROM event_discussions WHERE event_id = $1`,
		`DELETE FROM event_speakers WHERE event_id = $1`,
		`DELETE FROM event_category_mappings WHERE event_id = $1`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(ctx, query, eventID); err != nil {
			return fmt.Errorf(
				"failed to delete related event data: %w",
				err,
			)
		}
	}

	// ✅ TAMBAH: Hapus event setelah data relasi
	tag, err := tx.Exec(ctx, `
		DELETE FROM events
		WHERE event_id = $1
		  AND user_id = $2
	`, eventID, userID)

	if err != nil {
		return err
	}

	if tag.RowsAffected() != 1 {
		return ErrDeleteEventNotFound
	}

	// ✅ TAMBAH: Simpan perubahan transaksi
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// ✅ TAMBAH: Bersihkan cache Redis setelah commit
	if r.Redis != nil {
		redisCtx, cancel := context.WithTimeout(
			context.Background(),
			3*time.Second,
		)
		defer cancel()

		keys := []string{
			fmt.Sprintf(
				"eventhub:chace:event:detail:%d",
				eventID,
			),
			fmt.Sprintf("event_detail:%d", eventID),
			fmt.Sprintf("event:%d", eventID),
		}

		if err := r.Redis.Del(redisCtx, keys...).Err(); err != nil {
			log.Printf(
				"failed to invalidate deleted event cache %d: %v",
				eventID,
				err,
			)
		}
	}

	return nil
}
