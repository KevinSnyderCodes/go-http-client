package http

import "fmt"

// StatusCodeError represents an unexpected HTTP status code.
type StatusCodeError struct {
	StatusCode int
}

func (o *StatusCodeError) Error() string {
	return fmt.Sprintf("received status code %d", o.StatusCode)
}
