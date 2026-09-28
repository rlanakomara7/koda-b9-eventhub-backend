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
