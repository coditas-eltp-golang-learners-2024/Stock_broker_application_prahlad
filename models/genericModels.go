package models

// SuccessResponse represents a successful response format.
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response format.
type ErrorResponse struct {
	Error string `json:"error"`
}
