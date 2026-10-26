package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (s *EventService) GetEventsPaginated(
	ctx context.Context,
	search string,
	format string,
	category string,
	page int,
	limit int,
	path string,
) (dto.EventPaginationResponse, error) {
	result := dto.EventPaginationResponse{
		Data: make([]dto.EventCardResponse, 0),
	}

	total, err := s.EventRepo.CountEventsPaginated(
		ctx,
		search,
		format,
		category,
	)

	if err != nil {
		return result, err
	}

	events, err := s.EventRepo.GetEventsPage(
		ctx,
		search,
		format,
		category,
		page,
		limit,
	)

	if err != nil {
		return result, err
	}

	totalPage := (total + limit - 1) / limit

	result.Data = events
	result.Metadata = dto.PaginationMeta{
		CurrentPage: page,
		TotalPage:   totalPage,
		TotalData:   total,
		Limit:       limit,
		Next:        "",
		Prev:        "",
	}

	makeURL := func(targetPage int) string {
		params := url.Values{}

		params.Set("page", strconv.Itoa(targetPage))
		params.Set("limit", strconv.Itoa(limit))

		if search != "" {
			params.Set("search", search)
		}

		if format != "" {
			params.Set("format", format)
		}

		if category != "" {
			params.Set("category", category)
		}

		return path + "?" + params.Encode()
	}

	if page < totalPage {
		result.Metadata.Next = makeURL(page + 1)
	}

	if page > 1 && page <= totalPage {
		result.Metadata.Prev = makeURL(page - 1)
	}

	return result, nil
}
