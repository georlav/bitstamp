package bitstamp

import "time"

type option func(*HTTPAPI)
type wsOption func(*WebsocketAPI)

// BaseURLOption changes the default HTTP API base url
func BaseURLOption(val string) option {
	return func(api *HTTPAPI) {
		api.baseURL = val
	}
}

// APIKeyOption set the api key
func APIKeyOption(val string) option {
	return func(api *HTTPAPI) {
		api.key = val
	}
}

// APISecretOption set the api key
func APISecretOption(val string) option {
	return func(api *HTTPAPI) {
		api.secret = val
	}
}

// EnableDebugOption enable debuging (disabled by default)
func EnableDebugOption() option {
	return func(api *HTTPAPI) {
		api.debug = true
	}
}

// ClientTimeoutOption change default http client timeout
func ClientTimeoutOption(seconds int64) func(api *HTTPAPI) {
	return func(api *HTTPAPI) {
		api.handle.Timeout = time.Second * time.Duration(seconds)
	}
}

// WSSetAddressOption changes the default websocket address
func SetWSAddressOption(val string) wsOption {
	return func(api *WebsocketAPI) {
		api.address = val
	}
}
