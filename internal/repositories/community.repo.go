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
		community_id,
		name,
		description,
		banner_url
	FROM communities
	WHERE status = 'active'
	`

	args := []interface{}{}

	if search != "" {
		args = append(
			args,
			"%"+search+"%",
		)

		query += `AND name ILIKE $1`
	}

	query += `ORDER BY name ASC`

	rows, err := r.DB.Query(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	communities := []dto.CommunityResponse{}

	for rows.Next() {

		var community dto.CommunityResponse

		var description *string
		var bannerURL *string

		err := rows.Scan(
			&community.CommunityID,
			&community.Name,
			&description,
			&bannerURL,
		)

		if err != nil {
			return nil, err
		}

		if description != nil {
			community.Description = *description
		}

		if bannerURL != nil {
			community.BannerURL = *bannerURL
		}

		categoryRows, err := r.DB.Query(
			context.Background(),
			`
			SELECT cc.category_name
			FROM community_categories cc
			JOIN community_category_mappings ccm
				ON cc.category_id = ccm.category_id
			WHERE ccm.community_id = $1
			`,
			community.CommunityID,
		)

		if err != nil {
			return nil, err
		}

		community.Categories = []string{}

		for categoryRows.Next() {

			var categoryName string

			err := categoryRows.Scan(
				&categoryName,
			)

			if err != nil {
				categoryRows.Close()
				return nil, err
			}

			community.Categories = append(
				community.Categories,
				categoryName,
			)
		}

		categoryRows.Close()

		if category != "" {

			found := false

			for _, categoryName := range community.Categories {
				if categoryName == category {
					found = true
					break
				}
			}

			if !found {
				continue
			}
		}

		err = r.DB.QueryRow(
			context.Background(),
			`
			SELECT COUNT(*)
			FROM community_members
			WHERE community_id = $1
			`,
			community.CommunityID,
		).Scan(
			&community.MemberCount,
		)

		if err != nil {
			return nil, err
		}

		err = r.DB.QueryRow(
			context.Background(),
			`
			SELECT COUNT(*)
			FROM events
			WHERE community_id = $1
			AND event_date >= CURRENT_DATE
			AND status != 'completed'
			`,
			community.CommunityID,
		).Scan(
			&community.UpcomingCount,
		)

		if err != nil {
			return nil, err
		}

		var joinedCount int

		err = r.DB.QueryRow(
			context.Background(),
			`
			SELECT COUNT(*)
			FROM community_members
			WHERE community_id = $1
			AND user_id = $2
			`,
			community.CommunityID,
			userID,
		).Scan(
			&joinedCount,
		)

		if err != nil {
			return nil, err
		}

		community.IsJoined = joinedCount > 0

		communities = append(
			communities,
			community,
		)
	}

	return communities, nil
}
