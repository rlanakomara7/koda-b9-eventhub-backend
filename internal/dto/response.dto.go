package dto

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
