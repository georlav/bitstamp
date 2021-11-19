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

func newError(message string, statusCode int) Error {
	return Error{
		Message:    message,
		StatusCode: statusCode,
	}
}

func newErrorFromResponse(resp *http.Response) Error {
	if resp == nil {
		return newError("unable to parse error from a nil response", 0)
	}

	// API does not return valid json objects on those cases, returns text/html
	if resp.Header.Get("Content-Type") != "application/json" || resp.StatusCode == http.StatusNotFound {
		return newError(http.StatusText(resp.StatusCode), resp.StatusCode)
	}

	var errResp GenericErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return newError(fmt.Sprintf("unable to parse error response, %s", err), resp.StatusCode)
	}

	message := http.StatusText(resp.StatusCode)
	switch {
	case errResp.Reason != nil:
		message = string(errResp.Reason)
	case len(errResp.Errors) > 0:
		message = string(errResp.Errors)
	case errResp.Error != "":
		message = errResp.Error
	}

	return newError(message, resp.StatusCode)
}

func (e Error) Error() string {
	return e.Message
}
