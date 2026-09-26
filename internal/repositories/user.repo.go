package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

//onstructor

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// create user
func (r *UserRepository) Creatre(user *models.User) error {
	query := `INSERT INTO users (role_id,
		name,
		email,
		password,
		avatar_url,
		job,
		location,
		bio,
		status)
		VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING user_id
		`
	err := r.DB.QueryRow(context.Background(), query,
		user.RoleID,
		user.Name,
		user.Email,
		user.Password,
		user.AvatarURL,
		user.Job,
		user.Location,
		user.Bio,
		user.Status,
	).Scan(&user.UserID)
	return err
}

// find user
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `SELECT 
		user_id,
		role_id,
		name,
		email,
		password,
		avatar_url,
		job,
		location,
		bio,
		status,
		created_at,
		updated_at
		
		FROM users 
		WHERE email=$1`

	user := &models.User{}

	err := r.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.UserID,
		&user.RoleID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.AvatarURL,
		&user.Job,
		&user.Location,
		&user.Bio,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
