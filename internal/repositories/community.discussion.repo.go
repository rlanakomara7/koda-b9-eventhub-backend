package repositories

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *CommunityMemberRepository) GetCommunityDiscussions(
	ctx context.Context,
	communityID uint,
) ([]dto.CommunityDiscussionResponse, error) {
	rows, err := r.DB.Query(
		ctx,
		`
		SELECT
			d.discussion_id,
			d.community_id,
			d.user_id,
			u.name,
			COALESCE(u.avatar_url, ''),
			d.message,
			d.created_at
		FROM community_discussions d
		JOIN users u ON u.user_id = d.user_id
		JOIN communities c ON c.community_id = d.community_id
		WHERE d.community_id = $1
		  AND c.status = 'active'
		ORDER BY d.created_at DESC, d.discussion_id DESC
		LIMIT 100
		`,
		communityID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discussions := []dto.CommunityDiscussionResponse{}

	for rows.Next() {
		var discussion dto.CommunityDiscussionResponse

		err := rows.Scan(
			&discussion.DiscussionID,
			&discussion.CommunityID,
			&discussion.UserID,
			&discussion.Name,
			&discussion.AvatarURL,
			&discussion.Message,
			&discussion.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		discussions = append(discussions, discussion)
	}

	return discussions, rows.Err()
}

func (r *CommunityMemberRepository) CreateCommunityDiscussion(
	ctx context.Context,
	communityID uint,
	userID uint,
	message string,
) (dto.CommunityDiscussionResponse, error) {
	var discussion dto.CommunityDiscussionResponse

	err := r.DB.QueryRow(
		ctx,
		`
		WITH inserted AS (
			INSERT INTO community_discussions (
				community_id, user_id, message
			)
			SELECT c.community_id, $2, $3
			FROM communities c
			WHERE c.community_id = $1
			  AND c.status = 'active'
			  AND EXISTS (
				  SELECT 1
				  FROM community_members cm
				  WHERE cm.community_id = c.community_id
				    AND cm.user_id = $2
			  )
			RETURNING
				discussion_id,
				community_id,
				user_id,
				message,
				created_at
		)
		SELECT
			d.discussion_id,
			d.community_id,
			d.user_id,
			u.name,
			COALESCE(u.avatar_url, ''),
			d.message,
			d.created_at
		FROM inserted d
		JOIN users u ON u.user_id = d.user_id
		`,
		communityID,
		userID,
		message,
	).Scan(
		&discussion.DiscussionID,
		&discussion.CommunityID,
		&discussion.UserID,
		&discussion.Name,
		&discussion.AvatarURL,
		&discussion.Message,
		&discussion.CreatedAt,
	)

	return discussion, err
}
