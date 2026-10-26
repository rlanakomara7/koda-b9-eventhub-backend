package dto

type PaginationMeta struct {
	CurrentPage int    `json:"current_page"`
	TotalPage   int    `json:"total_page"`
	TotalData   int    `json:"total_data"`
	Limit       int    `json:"limit"`
	Next        string `json:"next"`
	Prev        string `json:"prev"`
}

type EventPaginationResponse struct {
	Data     []EventCardResponse `json:"data"`
	Metadata PaginationMeta      `json:"metadata"`
}
