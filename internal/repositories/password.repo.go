package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PasswordRepository struct {
	DB *pgxpool.Pool
}

func NewPasswordRepository(db *pgxpool.Pool) *PasswordRepository {
	return &PasswordRepository{
		DB: db,
	}
}

func (r *PasswordRepository) CreateResetToken(userID uint, token string) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO password_resets
		(user_id, token, expired_at)
		VALUES
		($1,$2,now()+interval '15 minute')
		`,
		userID,
		token,
	)
	return err
}

func (r *PasswordRepository) FindResetToken(token string) (uint, error) {
	var userID uint

	err := r.DB.QueryRow(
		context.Background(),
		`
		SELECT user_id
		FROM password_resets
		WHERE token=$1
		AND expired_at > now()
		`, token,
	).Scan(&userID)

	return userID, err
}

// Reset pass DB transaction
func (r *PasswordRepository) ResetPasswordTransaction(
	userID uint,
	token string,
	password string,
) error {

	ctx := context.Background()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`
		UPDATE users
		SET password=$1,
		updated_at=now()
		WHERE user_id=$2`,
		password,
		userID,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		`
		DELETE FROM password_resets
		WHERE token=$1`,
		token,
	)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
