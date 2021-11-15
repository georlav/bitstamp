package bitstamp

import (
	"fmt"
)

type APIError struct {
	Message    string
	StatusCode int
}

func NewAPIError(message string, statusCode int) APIError {
	return APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e APIError) Error() string {
	return fmt.Sprintf("status: %d, error: %s", e.StatusCode, e.Message)
}
