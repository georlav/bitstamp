package bitstamp

import (
	"encoding/json"
	"fmt"
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

	// API does not return valid json objects on those cases, returns text/html
	if resp.Header.Get("Content-Type") != "application/json" || resp.StatusCode == http.StatusNotFound {
		return NewError(http.StatusText(resp.StatusCode), resp.StatusCode)
	}

	var errResp GenericErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return NewError(fmt.Sprintf("unable to parse error response, %s", err), resp.StatusCode)
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
