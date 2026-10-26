package repositories

import (
	"context"
)

type EventCategory struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (r *EventRepository) GetEventCategories(
	ctx context.Context,
) ([]EventCategory, error) {
	const query = `
		SELECT category_id, category_name
		FROM event_categories
		ORDER BY category_name ASC
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]EventCategory, 0)

	for rows.Next() {
		var category EventCategory

		err := rows.Scan(
			&category.CategoryID,
			&category.CategoryName,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
