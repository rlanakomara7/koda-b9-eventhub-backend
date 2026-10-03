package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

//constructor

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// create user
func (r *UserRepository) Create(user *models.User) error {
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

// get user by id
func (r *UserRepository) GetUserByID(userID uint) (*models.User, error) {

	query := `
	SELECT 
		user_id,
		role_id,
		name,
		email,
		avatar_url,
		job,
		location,
		bio,
		status,
		created_at,
		updated_at
	FROM users
	WHERE user_id = $1
	`

	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.AvatarURL,
		&user.Job,
		&user.Location,
		&user.Bio,
		&user.Status,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// getprofile
func (r *UserRepository) GetProfile(userID uint) (*dto.ProfileResponse, error) {

	query := `
	SELECT
		u.name,
		u.email,
		u.avatar_url,
		u.location,
		u.created_at,
		r.role_name,
		u.bio
	FROM users u
	JOIN roles r ON u.role_id = r.role_id
	WHERE u.user_id = $1
	`

	profile := &dto.ProfileResponse{}

	var createdAt time.Time

	// karena kolom ini pada ERD boleh NULL
	var avatarURL *string
	var location *string
	var bio *string

	err := r.DB.QueryRow(
		context.Background(),
		query,
		userID,
	).Scan(
		&profile.Name,
		&profile.Email,
		&avatarURL,
		&location,
		&createdAt,
		&profile.Role,
		&bio,
	)

	if err != nil {
		return nil, err
	}

	if avatarURL != nil {
		profile.AvatarURL = *avatarURL
	}

	if location != nil {
		profile.Location = *location
	}

	if bio != nil {
		profile.Bio = *bio
	}

	profile.JoinedAt = createdAt.Format("January 2006")

	queryEvent := `
	SELECT COUNT(*)
	FROM event_members
	WHERE user_id = $1
	`

	err = r.DB.QueryRow(
		context.Background(),
		queryEvent,
		userID,
	).Scan(
		&profile.Stats.Events,
	)

	if err != nil {
		return nil, err
	}

	queryCommunity := `
	SELECT COUNT(*)
	FROM community_members
	WHERE user_id = $1
	`

	err = r.DB.QueryRow(
		context.Background(),
		queryCommunity,
		userID,
	).Scan(
		&profile.Stats.Communities,
	)

	if err != nil {
		return nil, err
	}

	querySaved := `
	SELECT COUNT(*)
	FROM saved_events
	WHERE user_id = $1
	`

	err = r.DB.QueryRow(
		context.Background(),
		querySaved,
		userID,
	).Scan(
		&profile.Stats.Saved,
	)

	if err != nil {
		return nil, err
	}

	return profile, nil
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

// update password
func (r *UserRepository) UpdatePassword(userID uint, password string) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		UPDATE users
		SET password=$1,
		updated_at=now()
		WHERE user_id=$2
		`,
		password,
		userID,
	)
	return err
}
