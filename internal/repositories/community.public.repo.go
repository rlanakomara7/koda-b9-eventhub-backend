package repositories

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (r *CommunityRepository) GetPublicCommunities(
	ctx context.Context,
) ([]dto.CommunityResponse, error) {
	query := `
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
			) AS categories,

			(
				SELECT COUNT(*)::int
				FROM community_members cm
				WHERE cm.community_id = c.community_id
			) AS member_count,

			(
				SELECT COUNT(*)::int
				FROM events e
				WHERE e.community_id = c.community_id
				  AND e.event_date >= CURRENT_DATE
				  AND e.status = 'active'
			) AS upcoming_count

		FROM communities c
		WHERE c.status = 'active'
		ORDER BY member_count DESC, c.name ASC
	`

	rows, err := r.DB.Query(ctx, query)
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
		)

		if err != nil {
			return nil, err
		}

		community.IsJoined = false
		communities = append(communities, community)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return communities, nil
}
