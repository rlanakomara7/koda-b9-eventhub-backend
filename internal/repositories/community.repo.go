package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

type CommunityRepository struct {
	DB *pgxpool.Pool
}

func NewCommunityRepository(db *pgxpool.Pool) *CommunityRepository {
	return &CommunityRepository{
		DB: db,
	}
}

func (r *CommunityRepository) GetCommunities(
	search string,
	category string,
	userID uint,
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
            ) AS upcoming_count,

            EXISTS (
                SELECT 1
                FROM community_members cm
                WHERE cm.community_id = c.community_id
                  AND cm.user_id = $3
            ) AS is_joined

        FROM communities c

        WHERE c.status = 'active'

          AND (
              $1 = ''
              OR c.name ILIKE '%' || $1 || '%'
          )

          AND (
              $2 = ''
              OR EXISTS (
                  SELECT 1
                  FROM community_category_mappings ccm
                  JOIN community_categories cc
                      ON cc.category_id = ccm.category_id
                  WHERE ccm.community_id = c.community_id
                    AND LOWER(cc.category_name) = LOWER($2)
              )
          )

        ORDER BY c.name ASC
    `

	rows, err := r.DB.Query(
		context.Background(),
		query,
		search,
		category,
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
