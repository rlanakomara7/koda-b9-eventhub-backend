package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CommunityMemberProfile struct {
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Job       string `json:"job"`
}

type CommunityMemberRepository struct {
	DB *pgxpool.Pool
}

func NewCommunityMemberRepository(
	db *pgxpool.Pool,
) *CommunityMemberRepository {
	return &CommunityMemberRepository{
		DB: db,
	}
}

func (r *CommunityMemberRepository) IsMember(
	communityID uint,
	userID uint,
) (bool, error) {
	var count int

	err := r.DB.QueryRow(
		context.Background(),
		`
		SELECT COUNT(*)
		FROM community_members
		WHERE community_id = $1
		  AND user_id = $2
		`,
		communityID,
		userID,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *CommunityMemberRepository) JoinCommunity(
	communityID uint,
	userID uint,
) error {

	_, err := r.DB.Exec(
		context.Background(),
		`
        WITH joined AS (
            INSERT INTO community_members (
                community_id, user_id
            )
            VALUES ($1, $2)
            RETURNING community_id, user_id
        )
        INSERT INTO notifications (
            user_id, type, title, message, path
        )
        SELECT
            j.user_id,
            'community',
            'Joined community',
            'You joined ' || c.name || '.',
            '/communities/' || c.community_id::text
        FROM joined j
        JOIN communities c
          ON c.community_id = j.community_id
        `,
		communityID,
		userID,
	)

	return err
}

func (r *CommunityMemberRepository) LeaveCommunity(
	communityID uint,
	userID uint,
) error {
	_, err := r.DB.Exec(
		context.Background(),
		`
		DELETE FROM community_members
		WHERE community_id = $1
		  AND user_id = $2
		`,
		communityID,
		userID,
	)

	return err
}

func (r *CommunityMemberRepository) GetMembers(
	ctx context.Context,
	communityID uint,
) ([]CommunityMemberProfile, error) {
	rows, err := r.DB.Query(
		ctx,
		`
		SELECT
			u.user_id,
			u.name,
			COALESCE(u.avatar_url, ''),
			COALESCE(u.job, '')
		FROM community_members cm
		JOIN users u
			ON u.user_id = cm.user_id
		WHERE cm.community_id = $1
		ORDER BY cm.joined_at ASC, cm.member_id ASC
		`,
		communityID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]CommunityMemberProfile, 0)

	for rows.Next() {
		var member CommunityMemberProfile

		if err := rows.Scan(
			&member.UserID,
			&member.Name,
			&member.AvatarURL,
			&member.Job,
		); err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
