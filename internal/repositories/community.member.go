package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CommunityMemberRepository struct {
	DB *pgxpool.Pool
}

func NewCommunityMemberRepository(db *pgxpool.Pool) *CommunityMemberRepository {
	return &CommunityMemberRepository{
		DB: db,
	}
}

func (r *CommunityMemberRepository) IsMember(
	communityID uint, userID uint) (bool, error) {

	var count int

	err := r.DB.QueryRow(
		context.Background(),
		`
			SELECT COUNT(*)
			FROM community_members
			WHERE community_id=$1
			AND user_id=$1
			`,
		communityID,
		userID,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// join
func (r *CommunityMemberRepository) JoinCommunity(
	communityID uint, userID uint,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		INSERT INTO community_members
		(community_id, user_id)
		VALUES ($1, $2)
		`,
		communityID,
		userID,
	)
	return err
}

// leave
func (r *CommunityMemberRepository) LeaveCommunity(
	communityID uint,
	userID uint,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
		DELETE FROM community_members
		WHERE community_id=$1
		AND user_id=$2
		`,
		communityID,
		userID,
	)

	return err
}
