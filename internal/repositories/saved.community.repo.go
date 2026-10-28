package repositories

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *CommunityMemberRepository) SaveCommunity(communityID, userID uint) error {
	_, err := r.DB.Exec(
		context.Background(),
		`INSERT INTO saved_communities (user_id, community_id)
		 VALUES ($1, $2)
		 ON CONFLICT (user_id, community_id) DO NOTHING`,
		userID, communityID,
	)

	return err
}

func (r *CommunityMemberRepository) UnsaveCommunity(communityID, userID uint) error {
	_, err := r.DB.Exec(
		context.Background(),
		`DELETE FROM saved_communities
		 WHERE user_id = $1 AND community_id = $2`,
		userID, communityID,
	)

	return err
}

func (r *CommunityMemberRepository) GetSavedCommunities(userID uint) ([]dto.CommunityResponse, error) {
	rows, err := r.DB.Query(
		context.Background(),
		`
		SELECT
			c.community_id,
			c.name,
			COALESCE(c.description, ''),
			COALESCE(c.banner_url, ''),
			ARRAY(
				SELECT cc.category_name
				FROM community_category_mappings ccm
				JOIN community_categories cc
					ON cc.category_id = ccm.category_id
				WHERE ccm.community_id = c.community_id
				ORDER BY cc.category_name
			),
			(
				SELECT COUNT(*)::int
				FROM community_members cm
				WHERE cm.community_id = c.community_id
			),
			(
				SELECT COUNT(*)::int
				FROM events e
				WHERE e.community_id = c.community_id
				  AND e.event_date >= CURRENT_DATE
				  AND e.status = 'active'
			),
			EXISTS(
				SELECT 1
				FROM community_members cm
				WHERE cm.community_id = c.community_id
				  AND cm.user_id = $1
			)
		FROM saved_communities sc
		JOIN communities c
			ON c.community_id = sc.community_id
		WHERE sc.user_id = $1
		  AND c.status = 'active'
		ORDER BY sc.saved_community_id DESC
		`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	communities := []dto.CommunityResponse{}

	for rows.Next() {
		var community dto.CommunityResponse

		err := rows.Scan(
			&community.CommunityID,
			&community.Name,
			&community.Description,
			&community.BannerURL,
			&community.Categories,
			&community.MemberCount,
			&community.UpcomingCount,
			&community.IsJoined,
		)

		if err != nil {
			return nil, err
		}

		communities = append(communities, community)
	}

	return communities, rows.Err()
}
