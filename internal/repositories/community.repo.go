package repositories

import (
	"context"
	"fmt"

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
SELECT DISTINCT
    c.community_id,
    c.name,
    COALESCE(c.description, ''),
    COALESCE(c.banner_url, ''),
    cm.user_id IS NOT NULL AS is_joined
FROM communities c
LEFT JOIN community_members cm
    ON cm.community_id = c.community_id
    AND cm.user_id = $3
LEFT JOIN community_category_mappings ccm
    ON ccm.community_id = c.community_id
LEFT JOIN community_categories cc
    ON cc.category_id = ccm.category_id
WHERE c.status = 'active'
AND (
    $1 = ''
    OR c.name ILIKE '%' || $1 || '%'
)
AND (
    $2 = ''
    OR cc.category_name ILIKE '%' || $2 || '%'
)
ORDER BY c.name ASC
`

	fmt.Println("QUERY COMMUNITY:")
	fmt.Println(query)

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
			&community.IsJoined,
		)

		if err != nil {
			return nil, err
		}

		communities = append(
			communities,
			community,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return communities, nil
}
