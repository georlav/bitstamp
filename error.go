package bitstamp

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func NewError(message string, statusCode int) Error {
	return Error{
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewErrorFromResponse(resp *http.Response) Error {
	if resp == nil {
		return NewError("unable to parse error from a nil response", 0)
	}

	var errResp GenericErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return NewError("unable to parse error response, %s", resp.StatusCode)
	}

	message := http.StatusText(resp.StatusCode)
	switch {
	case errResp.Reason != "":
		message = errResp.Reason
	case len(errResp.Errors) > 0:
		message = errResp.Errors[0].Message
	case errResp.Error != "":
		message = errResp.Error
	}

	return NewError(message, resp.StatusCode)
}

func (e Error) Error() string {
	return e.Message
}
